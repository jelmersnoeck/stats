package statsd_test

import (
	"testing"

	"github.com/jelmersnoeck/stats"
	"github.com/jelmersnoeck/stats/clients/statsd"
)

func BenchmarkStatsdClient(b *testing.B) {
	b.ReportAllocs()

	c, _ := statsd.New()
	stats := stats.AppStat{}

	for i := 0; i < b.N; i++ {
		c.Collect(stats)
	}
}
