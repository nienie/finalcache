package stats

import (
	"github.com/rcrowley/go-metrics"
)

//Counter ...
type Counter struct {
	counters map[string]int
}

//NewCounter ...
func NewCounter() *Counter {
	return &Counter{
		counters: make(map[string]int),
	}
}

//Incr ...
func (o *Counter) Incr(key string, cnt int) {
	o.counters[key] += cnt
}

//GetCounter ...
func (o *Counter) GetCounter(key string) int {
	return o.counters[key]
}

//Histogram ...
func (o *Counter) Histogram() metrics.Histogram {
	sample := metrics.NewUniformSample(len(o.counters))
	for _, cnt := range o.counters {
		sample.Update(int64(cnt))
	}
	return metrics.NewHistogram(sample)
}