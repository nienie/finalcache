package stats

import (
	"sync"
	"time"

	"github.com/nienie/window"
)

type access struct {
	count int
	t     int64 //unit: time.Duration
	key   string
}

//Stats ...
type Stats struct {
	mu             sync.Mutex
	percentile     float64
	windowSlide    time.Duration                 //滑动窗口滑动距离
	windowSize     time.Duration                 //滑动窗口大小
	windowAssigner *window.SlidingWindowAssigner //滑动窗口Assigner
	activeWindows  window.Windows                //当前活跃的窗口
	activeCounters CounterMap                    //当前活跃的Counter
	counter        *Counter                      //当前访问计数
	distribution   Distribution                  //当前访问的分布情况
	accessCh       chan access                   //访问计数通道
}

//NewStats ...
func NewStats(windowSize, windowSlide time.Duration, percentile float64) *Stats {
	stats := &Stats{
		percentile:     percentile,
		windowSize:     windowSize,
		windowSlide:    windowSlide,
		windowAssigner: window.NewSlidingWindowAssigner("finalcache", windowSize, windowSlide, -8*time.Hour),
		activeWindows:  window.NewWindows(),
		activeCounters: NewCounterMap(),
		accessCh:       make(chan access, 1024),
	}
	go stats.accessStatsLoop()
	go stats.triggerLoop()
	return stats
}

//Access ...
func (o *Stats) Access(key string, count int) {
	o.accessCh <- access{
		key:   key,
		t:     time.Now().UnixNano(),
		count: count,
	}
}

func (o *Stats) accessStatsLoop() {
	for ac := range o.accessCh {
		wnds := o.windowAssigner.AssignWindows(time.Duration(ac.t))
		o.mu.Lock()
		for _, w := range wnds {
			o.activeWindows.Add(w)
			counter := o.activeCounters.GetOrCreateCounter(w.GetName())
			counter.Incr(ac.key, ac.count)
		}
		o.mu.Unlock()
	}
}

//获取到期窗口，然后计算统计结果
func (o *Stats) triggerWindow() {
	for {
		w := o.activeWindows.Peek()
		if w == nil {
			return
		}
		//窗口还没过期，不需要计算
		now := time.Now()
		if time.Duration(now.UnixNano())-(w.GetEnd()) < 0 {
			return
		}
		//TODO: Add Log

		//窗口已经过期，触发计算
		//window从当前正在统计的windows队列弹出
		o.activeWindows.PopFront()
		counter := o.activeCounters.RemoveCounter(w.GetName())
		if counter != nil {
			o.distribution = NewDistribution(counter, o.percentile)
			o.counter = counter
		}
	}
}

func (o *Stats) triggerLoop() {
	var timer *time.Timer
	for {
		o.mu.Lock()
		o.triggerWindow()
		w := o.activeWindows.Peek()
		o.mu.Unlock()
		//目前没有窗口在计算，则等待o.windowSlide时长
		if w == nil {
			timer = time.NewTimer(o.windowSlide)
		} else {
			now := time.Now()
			//窗口过期了，需要重新计算窗口
			if time.Duration(now.UnixNano()) >= w.GetEnd() {
				continue
			}
			//等待窗口过期，往后推迟100ms
			timer = time.NewTimer(w.GetEnd() - time.Duration(now.UnixNano()) + 100*time.Millisecond)
		}
		//等待timer时间到期，到期后就可以触发窗口计算了
		select {
		case <-timer.C:
			timer.Stop()
		}
	}
}

//GetAccessCount ...
func (o *Stats) GetAccessCount(key string) int {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.counter == nil {
		return 0
	}
	return o.counter.GetCounter(key)
}