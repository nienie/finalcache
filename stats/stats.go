package stats

import "time"

type access struct {
	key string
	ts time.Time
	count int
}