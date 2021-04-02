package stats

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CounterTestSuite struct {
	suite.Suite
}

func (o *CounterTestSuite) TestAll() {
	counter := NewCounter()
	//1. 获取不存在的key
	noExistKey := "no-exist"
	count := counter.GetCounter(noExistKey)
	o.Zero(count)

	//2. key
	key1 := "key1"
	cnt := 1
	counter.Incr(key1, cnt)
	count = counter.GetCounter(key1)
	o.Equal(1, count)
	counter.Incr(key1, cnt)
	count = counter.GetCounter(key1)
	o.Equal(2, count)
	o.Equal(1, counter.GetKeyCount())

	key2 := "key2"
	counter.Incr(key2, cnt)
	o.Equal(2, counter.GetKeyCount())

	counter.Clear()
	o.Zero(counter.GetKeyCount())
}

func TestCounter(t *testing.T) {
	suite.Run(t, new(CounterTestSuite))
}

type CounterMapTestSuite struct {
	suite.Suite
}

func (o *CounterMapTestSuite) TestAll() {
	counterMap := NewCounterMap()
	var counter *Counter

	name1 := "counter1"
	counter1 := counterMap.GetOrCreateCounter(name1)
	o.NotNil(counter1)
	o.Equal(1, counterMap.Count())
	key1 := "key1"
	cnt := 2
	counter1.Incr(key1, cnt)

	counter = counterMap.GetOrCreateCounter(name1)
	o.Equal(counter1, counter)
	o.Equal(cnt, counter1.GetCounter(key1))

	name2 := "counter2"
	counter2 := counterMap.GetOrCreateCounter(name2)
	o.NotNil(counter2)
	o.Equal(2, counterMap.Count())

	counter = counterMap.RemoveCounter(name1)
	o.Equal(counter1, counter)
	o.Equal(1, counterMap.Count())

	counter = counterMap.GetOrCreateCounter(name2)
	o.Equal(counter2, counter)
}

func TestCounterMap(t *testing.T) {
	suite.Run(t, new(CounterMapTestSuite))
}