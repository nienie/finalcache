package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

)

type StatsTestSuite struct {
	suite.Suite
}

func (o *StatsTestSuite) TestAll() {
	stats := NewStats(3 * time.Second, 1 * time.Second, 0.999)
	key := "key"
	key1 := "key1"
	cnt := 6
	stats.Access(key, cnt)
	o.Equal(0, stats.GetAccessCount(key))

	stats.Access(key1, cnt)
	time.Sleep(1120 * time.Millisecond)
	o.Equal(cnt, stats.GetAccessCount(key))

	stats.Access(key1, cnt)
	time.Sleep(1120 * time.Millisecond)
	o.Equal(cnt, stats.GetAccessCount(key))

	stats.Access(key1, cnt)
	time.Sleep(1120 * time.Millisecond)
	o.Equal(0, stats.GetAccessCount(key))
}

func TestStats(t *testing.T) {
	suite.Run(t, new(StatsTestSuite))
}
