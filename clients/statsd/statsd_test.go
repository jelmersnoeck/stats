package statsd_test

import (
	"testing"

	"github.com/jelmersnoeck/stats"
	"github.com/jelmersnoeck/stats/clients/statsd"
	sd "gopkg.in/alexcesaro/statsd.v2"
)

func BenchmarkStatsdClient(b *testing.B) {
	b.ReportAllocs()

	c, err := statsd.New(sd.Mute(true))
	if err != nil {
		b.Error(err)
	}

	stats := stats.AppStat{}

	for i := 0; i < b.N; i++ {
		c.Collect(stats)
	}
}
