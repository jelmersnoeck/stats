package stats

import (
	"testing"
	"time"
)

func TestDefaultOptions(t *testing.T) {
	ops := newOptions()

	if ops.interval != time.Minute {
		t.Errorf("Expected default interval to be one minute")
	}

	if ops.client == nil {
		t.Errorf("Expected default client to be set")
	}

	if !ops.collectMemory {
		t.Errorf("Expected memory collection to be enabled by default")
	}

	if !ops.collectGoroutines {
		t.Errorf("Expected goroutine collection to be enabled by default")
	}

	if !ops.collectGC {
		t.Errorf("Expected Garbage Collection collection to be enabled by default")
	}

	if !ops.collectCgo {
		t.Errorf("Expected Cgo collection to be enabled by default")
	}
}

func TestCollectionInterval(t *testing.T) {
	interval := 15 * time.Second
	col := New(CollectionInterval(interval))

	if col.options.interval != interval {
		t.Errorf("Expected interval to be set to options value")
	}
}

func TestCollectionClient(t *testing.T) {
	client := &mockClient{}
	col := New(CollectionClient(client))

	if col.options.client != client {
		t.Errorf("Expected client to be set to options value")
	}
}

func TestCollectMemory(t *testing.T) {
	if New(CollectMemory(false)).options.collectMemory {
		t.Errorf("Expected memory collection to be disabled")
	}
}

func TestCollectGoroutines(t *testing.T) {
	if New(CollectGoroutines(false)).options.collectGoroutines {
		t.Errorf("Expected goroutine collection to be disabled")
	}
}

func TestCollectGC(t *testing.T) {
	if New(CollectGC(false)).options.collectGC {
		t.Errorf("Expected GC collection to be disabled")
	}
}

func TestCollectCgo(t *testing.T) {
	if New(CollectCgo(false)).options.collectCgo {
		t.Errorf("Expected Cgo collection to be disabled")
	}
}

type mockClient struct{}

func (c *mockClient) Collect(AppStat) {}
