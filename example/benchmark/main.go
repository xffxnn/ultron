package main

import (
	"context"
	"time"

	"github.com/wosai/ultron/v2"
)

type benchmarkAttacker struct {
	name string
}

func (b *benchmarkAttacker) Name() string {
	return b.name
}

func (b *benchmarkAttacker) Fire(_ context.Context) error {
	time.Sleep(10 * time.Millisecond)
	return nil
}

func main() {
	runner := ultron.NewLocalRunner()
	task := ultron.NewTask()
	task.Add(&benchmarkAttacker{name: "benchmark"}, 1)
	runner.Assign(task)

	plan := ultron.NewPlan("benchmark test")
	plan.AddStages(
		&ultron.V1StageConfig{ConcurrentUsers: 200, Duration: 30 * time.Second, RampUpPeriod: 10},
		// &ultron.V1StageConfig{ConcurrentUsers: 100, Requests: 500000, RampUpPeriod: 5},
	)

	if err := runner.Launch(); err != nil {
		panic(err)
	}

	if err := runner.StartPlan(plan); err != nil {
		panic(err)
	}

	select {}
}
