package statistics

import (
	"errors"
	"sort"
	"strconv"
	"sync"
	"time"
)

var (
	timeDistributions = []float64{0.5, 0.6, 0.7, 0.8, 0.9, 0.95, 0.97, 0.98, 0.99, 1.0}
)

const (
	CurrentTPSTimeRange = 12 * time.Second
	Total               = "[TOTAL]"
)

type (
	// AttackResut 事务执行结果
	AttackResut struct {
		Name     string
		Duration time.Duration
		Error    error
	}

	AttackStatistician struct {
		name                string                   // 事务名称
		requests            uint64                   // 成功请求数
		failures            uint64                   // 失败请求数
		totalResponseTime   time.Duration            // 原始响应时间汇总
		minResponseTime     time.Duration            // 最小响应时间
		maxResponseTime     time.Duration            // 最长响应时间
		recentSuccessBucket *timeRangeContainer      // 最近的成功请求数量
		recentFailureBucket *timeRangeContainer      // 最近的失败请求数量
		responseBucket      map[time.Duration]uint64 // 成功请求的响应时间桶
		failureBucket       map[string]uint64        // 失败请求的错误原因桶
		firstAttack         time.Time                // 请求开始时间
		lastAttack          time.Time                // 最后一次收到响应结果的时间
		interval            time.Duration            // 统计CurrentTPS（）的时间区间
		mu                  sync.RWMutex
	}

	// AttackReport 聚合报告
	AttackReport struct {
		Name           string                   // 事务名称
		Requests       uint64                   // 成功请求总数
		Failures       uint64                   // 失败请求总数
		Min            time.Duration            // 最小延迟
		Max            time.Duration            // 最大延迟
		Median         time.Duration            // 中位数
		Average        time.Duration            // 平均数
		TPS            float64                  // 每秒事务数
		Distributions  map[string]time.Duration // 百分位分布
		FailRation     float64                  // 错误率
		FailureDetails map[string]int32         // 错误详情分布
		FullHistory    bool                     // 是否是该阶段完整的报告
	}

	SummaryReport struct {
		PlanID  string
		Reports map[string]*AttackReport
	}

	timeRangeContainer struct {
		container map[int64]int64
		timeRange int64
	}

	StatisticianGroup struct {
		planID    string
		container map[string]*AttackStatistician // 优于sync.Map
		mu        sync.Mutex                     // 写多读少场景，互斥锁更好
	}
)

func newTimeRangeContainer(n int64) *timeRangeContainer {
	return &timeRangeContainer{
		container: make(map[int64]int64),
		timeRange: n,
	}
}

func (ls *timeRangeContainer) accumulate(k, v int64) {
	if _, ok := ls.container[k]; ok {
		ls.container[k] += v
		return
	}

	for key := range ls.container {
		if (k - key) > ls.timeRange {
			delete(ls.container, key)
		}
	}
	ls.container[k] = v
}

func findReponseBucket(t time.Duration) time.Duration {
	if t <= 100*time.Millisecond {
		return t
	}
	if t <= 1000*time.Millisecond {
		return (t + 5*time.Millisecond) / 1e7 * 1e7
	}
	return (t + 50*time.Millisecond) / 1e8 * 1e8
}

// IsFailure 事务是否执行失败
func (ar *AttackResut) IsFailure() bool {
	return ar.Error != nil
}

func NewAttackStatistician(name string) *AttackStatistician {
	if name == Total {
		panic("attacker name conflicts with build-in name")
	}
	return &AttackStatistician{
		name:                name,
		recentSuccessBucket: newTimeRangeContainer(20),
		recentFailureBucket: newTimeRangeContainer(20),
		responseBucket:      make(map[time.Duration]uint64),
		failureBucket:       make(map[string]uint64),
		interval:            CurrentTPSTimeRange,
	}
}

func (ara *AttackStatistician) recordSuccess(ret *AttackResut) {
	if ara.name != ret.Name {
		return
	}
	ara.mu.Lock()
	defer ara.mu.Unlock()

	ara.requests++
	ara.totalResponseTime += ret.Duration

	now := time.Now()
	if ara.firstAttack.IsZero() { // 第一次记录，且是成功请求
		ara.firstAttack = now
		ara.minResponseTime = ret.Duration
		ara.maxResponseTime = ret.Duration
	} else if ara.minResponseTime == 0 && ara.maxResponseTime == 0 { // 如果第一次先到的是错误请求，则min repsone 必然为0
		ara.minResponseTime = ret.Duration
		ara.maxResponseTime = ret.Duration
	} else {
		if ret.Duration < ara.minResponseTime {
			ara.minResponseTime = ret.Duration
		}
		if ret.Duration > ara.maxResponseTime {
			ara.maxResponseTime = ret.Duration
		}
	}
	ara.lastAttack = now

	ara.recentSuccessBucket.accumulate(now.Unix(), 1)
	ara.responseBucket[findReponseBucket(ret.Duration)]++
}

func (ara *AttackStatistician) recordFailure(ret *AttackResut) {
	if ara.name != ret.Name {
		return
	}

	ara.mu.Lock()
	defer ara.mu.Unlock()

	ara.failures++

	now := time.Now()
	if ara.firstAttack.IsZero() {
		ara.firstAttack = now
	}
	ara.lastAttack = now

	ara.failureBucket[ret.Error.Error()]++
	ara.recentFailureBucket.accumulate(now.Unix(), 1)
}

func (ara *AttackStatistician) Record(ret *AttackResut) {
	if ret.Error == nil {
		ara.recordSuccess(ret)
		return
	}
	ara.recordFailure(ret)
}

// TotalTPS 全程TPS
func (ara *AttackStatistician) totalTPS() float64 {
	if ara.lastAttack == ara.firstAttack {
		return 0
	}
	return float64(ara.requests) / float64(ara.lastAttack.Sub(ara.firstAttack).Seconds())
}

// CurrentTPS 最近12秒的TPS
func (ara *AttackStatistician) currentTPS() float64 {
	if ara.lastAttack == ara.firstAttack {
		return 0
	}

	end := time.Now().Add(-1 * time.Second) // 当前一秒未完成，往前推一秒作为统计终点
	if end.Before(ara.lastAttack) {         // 尚未执行满一秒，不统计
		return 0
	}
	start := end.Add(-1 * (ara.interval - time.Second))
	if start.Before(ara.firstAttack) {
		start = ara.firstAttack
	}
	if end.Sub(start) <= 0 {
		return 0
	}

	var count int64
	for i := start.Unix(); i <= end.Unix(); i++ {
		if v, ok := ara.recentSuccessBucket.container[i]; ok {
			count += v
		}
	}
	return float64(count) / float64(end.Unix()-start.Unix()+1)
}

func (ara *AttackStatistician) percentile(ps ...float64) []time.Duration {
	var bucketKeys []time.Duration
	for k := range ara.responseBucket {
		bucketKeys = append(bucketKeys, k)
	}
	sort.Slice(bucketKeys, func(i, j int) bool {
		return bucketKeys[i] < bucketKeys[j]
	})

	results := make([]time.Duration, len(ps))

percent:
	for n, per := range ps {
		index := int64(float64(ara.requests)*per + .5)
		if index >= int64(ara.requests) {
			results[n] = ara.maxResponseTime
			continue percent
		}
		if index <= 1 {
			results[n] = ara.minResponseTime
			continue percent
		}

		for _, key := range bucketKeys {
			index -= int64(ara.responseBucket[key])
			if index <= 0 {
				results[n] = key
				continue percent
			}
		}
		panic("unreachable code")
	}
	return results
}

func (ara *AttackStatistician) min() time.Duration {
	return ara.minResponseTime
}

func (ara *AttackStatistician) max() time.Duration {
	return ara.maxResponseTime
}

func (ara *AttackStatistician) average() time.Duration {
	if ara.requests == 0 {
		return 0
	}
	return time.Duration(ara.totalResponseTime / time.Duration(ara.requests))
}

func (ara *AttackStatistician) failRatio() float64 {
	total := float64(ara.requests) + float64(ara.failures)
	if total == 0 {
		return 0.0
	}
	return float64(ara.failures) / total
}

func (ara *AttackStatistician) Report(full bool) *AttackReport {
	ara.mu.RLock()
	defer ara.mu.RUnlock()

	report := &AttackReport{
		Name:           ara.name,
		Requests:       ara.requests,
		Failures:       ara.failures,
		Min:            ara.min(),
		Max:            ara.max(),
		Average:        ara.average(),
		Distributions:  make(map[string]time.Duration),
		FailRation:     ara.failRatio(),
		FailureDetails: make(map[string]int32),
		FullHistory:    full,
	}
	if full {
		report.TPS = ara.totalTPS()
	} else {
		report.TPS = ara.currentTPS()
	}
	pers := ara.percentile(timeDistributions...)
	for index, d := range timeDistributions {
		report.Distributions[strconv.FormatFloat(d, 'f', 2, 64)] = pers[index]
	}
	report.Median = pers[0]
	return report
}

func (ara *AttackStatistician) merge(other *AttackStatistician) error {
	if other == nil {
		return nil
	}
	if ara.name != other.name {
		return errors.New("cannot merge two different types report")
	}

	ara.mu.Lock()
	defer ara.mu.Unlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	ara.requests += other.requests
	ara.failures += other.failures
	ara.totalResponseTime += other.totalResponseTime
	if other.minResponseTime < ara.minResponseTime && other.minResponseTime > 0 {
		ara.minResponseTime = other.minResponseTime
	}
	if other.maxResponseTime > ara.maxResponseTime {
		ara.maxResponseTime = other.maxResponseTime
	}
	for k, v := range other.recentSuccessBucket.container {
		ara.recentSuccessBucket.accumulate(k, v)
	}
	for k, v := range other.recentFailureBucket.container {
		ara.recentFailureBucket.accumulate(k, v)
	}
	for k, v := range other.responseBucket {
		ara.responseBucket[k] += v
	}
	for k, v := range other.failureBucket {
		ara.failureBucket[k] += v
	}
	if other.firstAttack.Before(ara.firstAttack) {
		ara.firstAttack = other.firstAttack
	}
	if ara.lastAttack.Before(other.lastAttack) {
		ara.lastAttack = other.lastAttack
	}
	return nil
}

func (ara *AttackStatistician) BatchMerge(others ...*AttackStatistician) error {
	for _, other := range others {
		if err := ara.merge(other); err != nil {
			return err
		}
	}
	return nil
}

func NewStatistician() *StatisticianGroup {
	return &StatisticianGroup{container: make(map[string]*AttackStatistician)}
}

func (s *StatisticianGroup) Report(full bool) *SummaryReport {
	sr := &SummaryReport{PlanID: s.planID, Reports: make(map[string]*AttackReport)}
	sr.Reports[Total] = &AttackReport{Name: Total}

	s.mu.Lock()
	defer s.mu.Unlock()

	for key, value := range s.container {
		sr.Reports[key] = value.Report(full)
		sr.Reports[Total].Requests += sr.Reports[key].Requests
		sr.Reports[Total].Failures += sr.Reports[key].Failures
		sr.Reports[Total].TPS += sr.Reports[key].TPS
	}
	return sr
}

func (s *StatisticianGroup) Record(result *AttackResut) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if agg, ok := s.container[result.Name]; !ok {
		agg = NewAttackStatistician(result.Name)
		agg.Record(result)
		s.container[result.Name] = agg
	} else {
		agg.Record(result)
	}
}

func (s *StatisticianGroup) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key := range s.container {
		delete(s.container, key)
	}
}

func (s *StatisticianGroup) ReplaceStatistician(agg *AttackStatistician) error {
	if agg == nil {
		return errors.New("cannot replace with nil pointer")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.container[agg.name] = agg
	return nil
}
