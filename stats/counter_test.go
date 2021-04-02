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
	o.Require().Zero(count)

	//2. key
	key1 := "key1"
	cnt := 1
	counter.Incr(key1, cnt)
	count = counter.GetCounter(key1)
	o.Require().Equal(1, count)
	counter.Incr(key1, cnt)
	count = counter.GetCounter(key1)
	o.Require().Equal(2, count)
}

func TestCounter(t *testing.T) {
	suite.Run(t, new(CounterTestSuite))
}
