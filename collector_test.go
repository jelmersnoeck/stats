package stats

import (
	"encoding/json"
	"runtime"
	"sync"
	"testing"
)

func TestCapture(t *testing.T) {
	runtime.GC()

	var wg sync.WaitGroup
	col := &testCollector{wg: &wg}
	cl := New(CollectionClient(col))

	cl.capture()
	runData(&wg)

	if col.Memory <= 0 {
		t.Errorf("Expected memory to be captured")
	}

	if col.Goroutines <= 0 {
		t.Errorf("Exected a goroutine to be created")
	}

	if col.NumGC <= 0 {
		t.Errorf("Exected a GC to have run")
	}

	if len(col.GCPauses) <= 0 {
		t.Errorf("Exected a GC Pause to have been recorded")
	}

	if col.NumCgo <= 0 {
		// TODO: Need to figure out a way to test this with stdlib
		//t.Errorf("Expected CGO calls to have happened")
	}
}

func BenchmarkCapture(b *testing.B) {
	b.ReportAllocs()

	c := New()
	for i := 0; i < b.N; i++ {
		c.capture()
	}
}

// runData generates some information for all the stats that we are trying to
// collect. This assures that we can run our tests against something.
func runData(wg *sync.WaitGroup) {
	// WaitGroup to initialise the data run
	wg.Add(1)

	// allocate memory
	bts := []byte(`{"name": "Jelmer Snoeck"}`)
	var tp struct {
		Name string `json:"name"`
	}
	json.Unmarshal(bts, &tp)

	// get goroutine counter
	wg.Add(1)
	go func(g *sync.WaitGroup) {
		defer g.Done()
	}(wg)

	wg.Wait()
}

type testCollector struct {
	AppStat
	wg *sync.WaitGroup
}

func (c *testCollector) Collect(a AppStat) {
	c.AppStat = a
	c.wg.Done()
}
