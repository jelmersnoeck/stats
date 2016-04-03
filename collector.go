package stats

import (
	"log"
	"runtime"
	"runtime/debug"
	"time"
)

type collector struct {
	options collectionOptions
	lastGC  int64
}

// New creates a new client that is set up to collect application statistics.
func New(opts ...collectionOption) *collector {
	return &collector{
		options: newOptions(opts...),
	}
}

// Collect will run indefinitely at every interval configured by the options
// given to the collector and collects all the enabled metrics. These metrics
// will then be send to the specified client. It is suggested to run this in a
// goroutine.
func (c *collector) Collect() {
	defer recoverCheck()

	ticker := time.Tick(c.options.interval)
	for range ticker {
		go c.capture()
	}
}

func (c *collector) capture() {
	defer recoverCheck()

	stats := AppStat{}

	if c.options.collectMemory {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		stats.Memory = mem.Alloc
	}

	if c.options.collectGoroutines {
		stats.Goroutines = runtime.NumGoroutine()
	}

	if c.options.collectGC {
		var gcStats debug.GCStats
		debug.ReadGCStats(&gcStats)

		numGC := gcStats.NumGC
		if c.lastGC != numGC {
			stats.NumGC = numGC
			stats.GCPauses = gcStats.Pause[0:(numGC - c.lastGC)]
			c.lastGC = numGC
		}
	}

	if c.options.collectCgo {
		stats.NumCgo = runtime.NumCgoCall()
	}

	go func() {
		defer recoverCheck()
		c.options.client.Collect(stats)
	}()
}

func recoverCheck() {
	if rec := recover(); rec != nil {
		log.Println(rec)
	}
}
