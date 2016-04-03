package stats

// Client is used to send data to a storage.
type Client interface {
	Collect(AppStat)
}

// NilClient is a collector that gets data and disregards it immediately. It is
// the default collector. This collector should not be used for production as it
// will not collect any data.
type NilClient struct{}

func (c *NilClient) Collect(AppStat) {}
