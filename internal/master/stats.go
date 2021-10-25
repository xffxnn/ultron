package master

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/wosai/ultron/v2"
	"github.com/wosai/ultron/v2/pkg/statistics"
	"golang.org/x/sync/errgroup"
)

type (
	statsAggregator struct {
		counter           uint32
		toleranceForDelay uint32 // 容忍延后的批次
		providers         map[string]ultron.StatsProvider
		buffer            map[uint32]map[string]*statistics.StatisticianGroup
		mu                sync.Mutex
	}
)

var _ ultron.StatsAggregator = (*statsAggregator)(nil)

func newStatsAggregator() *statsAggregator {
	return &statsAggregator{
		providers:         make(map[string]ultron.StatsProvider),
		buffer:            make(map[uint32]map[string]*statistics.StatisticianGroup),
		toleranceForDelay: 3,
	}
}

func (agg *statsAggregator) Aggregate(fullHistory bool) (statistics.SummaryReport, error) {
	agg.mu.Lock()
	var providers []ultron.StatsProvider
	batch := agg.counter
	agg.counter++

	agg.buffer[batch] = make(map[string]*statistics.StatisticianGroup)
	for _, provider := range agg.providers {
		providers = append(providers, provider)
		agg.buffer[batch][provider.ID()] = nil
	}
	agg.mu.Unlock()

	if len(providers) == 0 {
		return statistics.SummaryReport{}, errors.New("no provider exists")
	}

	// 开始接收各个provider上报的数据
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	for _, provider := range providers {
		provider := provider // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			sg, err := provider.Submit(ctx, batch)
			if err != nil {
				return err
			}
			agg.mu.Lock()
			defer agg.mu.Unlock()
			agg.buffer[batch][provider.ID()] = sg
			return nil
		})
	}
	err := g.Wait()

	agg.mu.Lock()
	returns := agg.buffer[batch]
	delete(agg.buffer, batch) // 不管结果是否异常，都要从buffer中移除

	if err != nil {
		return statistics.SummaryReport{}, fmt.Errorf("batch-%d: %w", batch, err)
	}

	if (agg.counter - (batch + 1)) > agg.toleranceForDelay { // too late
		return statistics.SummaryReport{}, fmt.Errorf("batch-%d: too late to accept summary report", batch)
	}
	agg.mu.Unlock()

	// 检查是否完成
	sg := statistics.NewStatisticianGroup()
	for id, ret := range returns {
		if ret == nil {
			return statistics.SummaryReport{}, fmt.Errorf("batch-%d: provide %s submitted empty report", batch, id)
		} else {
			sg.Merge(ret)
		}
	}
	return sg.Report(fullHistory), nil
}

func (agg *statsAggregator) Remove(id string) {
	agg.mu.Lock()
	defer agg.mu.Unlock()

	delete(agg.providers, id)
}

func (agg *statsAggregator) Add(providers ...ultron.StatsProvider) {
	agg.mu.Lock()
	defer agg.mu.Unlock()

	for _, provider := range providers {
		agg.providers[provider.ID()] = provider
	}
}
