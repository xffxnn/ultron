package main

import (
	"context"
	"net/http"

	"github.com/wosai/ultron/handler/influxdbv1/v2"
	"github.com/wosai/ultron/v2"
)

func main() {
	runner := ultron.NewLocalRunner()

	// setup task
	task := ultron.NewTask()
	bing := ultron.NewHTTPAttacker("bing")
	bing.Apply(
		ultron.WithPrepareFunc(func(context.Context) (*http.Request, error) {
			return http.NewRequest(http.MethodGet, "https://bing.com", nil)
		}),
		ultron.WithCheckFuncs(ultron.CheckHTTPStatusCode),
	)
	task.Add(bing, 1)
	runner.Assign(task)

	// setup influxdb handler
	handler := influxdbv1.NewInfluxDBV1Handler()
	handler.Apply(
		influxdbv1.WithHTTPClient("127.0.0.1:8089", "", ""),
	)
	runner.SubscribeReport(handler.HandleReport())
	runner.SubscribeResult(handler.HandleResult(0.1))

	// start local runner
	if err := runner.Launch(); err != nil {
		panic(err)
	}

	// open http://localhost:2017
	select {}
}
