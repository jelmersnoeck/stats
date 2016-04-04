package stats

import "time"

type collectionOptions struct {
	interval time.Duration
	client   Client

	collectMemory     bool
	collectGoroutines bool
	collectGC         bool
	collectCgo        bool
}

type collectionOption func(*collectionOptions)

func newOptions(opts ...collectionOption) collectionOptions {
	options := collectionOptions{
		interval:          60 * time.Second,
		client:            &NilClient{},
		collectMemory:     true,
		collectGoroutines: true,
		collectGC:         true,
		collectCgo:        true,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// CollectionInterval sets the duration between collection runs. The default is
// 60 seconds.
func CollectionInterval(d time.Duration) collectionOption {
	return func(o *collectionOptions) {
		o.interval = d
	}
}

// CollectionClient represents the collection client we'll be using to send our
// collected data to. This can be an HTTP Client, a Statsd client, etc. as long
// as it conforms to the Client interface. The default client is the nil client,
// this client disregards any input that is passed through.
func CollectionClient(c Client) collectionOption {
	return func(o *collectionOptions) {
		o.client = c
	}
}

// CollectMemory is the option to enable or disable memory collection. This is
// enabled by default.
func CollectMemory(c bool) collectionOption {
	return func(o *collectionOptions) {
		o.collectMemory = c
	}
}

// CollectGoroutines is the option to enable or disable goroutine collection.
// This is enabled by default.
func CollectGoroutines(c bool) collectionOption {
	return func(o *collectionOptions) {
		o.collectGoroutines = c
	}
}

// CollectGC is the option to enable or disable collection of Garbage Collection
// information. This includes number of GC runs and the GC run times. This is
// enabled by default.
func CollectGC(c bool) collectionOption {
	return func(o *collectionOptions) {
		o.collectGC = c
	}
}

// CollectCgo is the option to enable or disable collection of the number of Cgo
// calls made by the application. This is enabled by default.
func CollectCgo(c bool) collectionOption {
	return func(o *collectionOptions) {
		o.collectCgo = c
	}
}
