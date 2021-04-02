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

//Incr 增加某个key的访问次数
func (o *Counter) Incr(key string, cnt int) {
	o.counters[key] += cnt
}

//GetCounter 获取某个key的访问次数
func (o *Counter) GetCounter(key string) int {
	return o.counters[key]
}

//GetKeyCount 获取key的数量
func (o *Counter) GetKeyCount() int {
	return len(o.counters)
}

//Clear ...
func (o *Counter) Clear() {
	o.counters = make(map[string]int)
}

//Histogram ...
func (o *Counter) Histogram() metrics.Histogram {
	sample := metrics.NewUniformSample(o.GetKeyCount())
	for _, cnt := range o.counters {
		sample.Update(int64(cnt))
	}
	return metrics.NewHistogram(sample)
}

//CounterMap ...
type CounterMap map[string]*Counter

//NewCounterMap ...
func NewCounterMap() CounterMap {
	return make(map[string]*Counter)
}

//GetOrCreateCounter ...
func (o *CounterMap) GetOrCreateCounter(name string) *Counter {
	counter, ok := (*o)[name]
	if !ok {
		counter = NewCounter()
		(*o)[name] = counter
	}
	return counter
}

//RemoveCounter ...
func (o *CounterMap) RemoveCounter(name string) *Counter {
	counter := (*o)[name]
	delete(*o, name)
	return counter
}

//Count ...
func (o *CounterMap) Count() int {
	return len(*o)
}
