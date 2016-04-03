package stats

import "time"

type AppStat struct {
	Memory     uint64
	Goroutines int
	Files      int
	NumGC      int64
	GCPauses   []time.Duration
	NumCgo     int64
}
