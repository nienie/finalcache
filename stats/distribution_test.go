package stats

import (
	"fmt"
	"testing"
)

func TestDistribution(t *testing.T) {
	counter := NewCounter()
	for i := 0; i < 10; i++ {
		counter.Incr(fmt.Sprint(i), i)
	}
	d := NewDistribution(counter, 0.99)
	t.Logf("distribution=%+v", d)
}