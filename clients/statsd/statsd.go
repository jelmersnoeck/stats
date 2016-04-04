package statsd

import (
	"time"

	"github.com/jelmersnoeck/stats"
	sd "gopkg.in/alexcesaro/statsd.v2"
)

// Client is an extension on the statsd client which implements a Collect method
// which is compatible with the collector.
type Client struct {
	*sd.Client
}

// New creates a new statsd compatible client for the stats collector.
func New(opts ...sd.Option) (*Client, error) {
	cl, err := sd.New(opts...)
	if err != nil {
		return nil, err
	}

	return &Client{Client: cl}, nil
}

// Collect will take an AppStat and put the data into statsd metrics. All the
// metrics will be prefixed with `appstat`.
func (s *Client) Collect(stats stats.AppStat) {
	s.Count("appstat.memory", stats.Memory)
	s.Count("appstat.goroutines", stats.Goroutines)
	s.Count("appstat.files", stats.Files)
	s.Count("appstat.num_gc", stats.NumGC)
	s.Count("appstat.num_cgo", stats.NumCgo)

	for _, d := range stats.GCPauses {
		s.Timing("appstat.gc_pause", int(d/time.Millisecond))
	}
}
