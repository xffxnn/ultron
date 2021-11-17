package ultron

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	chimiddleware "github.com/jacexh/gopkg/chi-middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type (
	restServer struct {
		runner *masterRunner
	}

	restResponse struct {
		Result       bool   `json:"result,omitempty"`
		ErrorMessage string `json:"error_message,omitempty"`
	}

	requestStartPlan struct {
		Name   string           `json:"name"`
		Stages []*V1StageConfig `json:"stages"`
	}
)

var (
	//go:embed web/static/*
	statics embed.FS
	//go:embed web/index.html
	indexhtml []byte
)

func (rest *restServer) handleStartNewPlan() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req := new(requestStartPlan)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			Logger.Error("failed to parse request body", zap.Error(err))
			renderResponse(err, rw, r)
			return
		}

		plan := NewPlan("")
		for _, stage := range req.Stages {
			plan.AddStages(stage)
		}
		err := rest.runner.StartPlan(plan)
		renderResponse(err, rw, r)
	}
}

func (rest *restServer) handleStopPlan() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rest.runner.StopPlan()
		renderResponse(nil, rw, r)
	}
}

func renderResponse(err error, w http.ResponseWriter, r *http.Request) {
	ret := &restResponse{}
	if err == nil {
		ret.Result = true
	} else {
		ret.ErrorMessage = err.Error()
	}

	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func buildHTTPRouter(runner *masterRunner) http.Handler {
	route := chi.NewRouter()
	rest := &restServer{runner: runner}

	{ // middlewares
		route.Use(middleware.RequestID)
		route.Use(middleware.RealIP)
		route.Use(chimiddleware.RequestZapLog(Logger))
		route.Use(middleware.Recoverer)
	}

	// http api
	{
		route.Post("/api/v1/plan", rest.handleStartNewPlan())
		route.Delete("/api/v1/plan", rest.handleStopPlan())
	}

	// static files
	content, err := fs.Sub(statics, "web/static")
	if err != nil {
		panic(err)
	}
	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(indexhtml)
	})
	route.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(content))))

	// prometheus exporter
	exporter := newMetric(runner)
	prometheus.MustRegister(exporter)
	runner.SubscribeReport(exporter.handleReport()) // 订阅report
	route.Handle("/metrics", promhttp.Handler())
	return route
}
