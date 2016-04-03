# Stats

[![Build Status](https://travis-ci.org/jelmersnoeck/stats.svg?branch=master)](https://travis-ci.org/jelmersnoeck/stats)

Stats is a statistics collection library for your Go applications. It collects
data on a regular, configurable interval and sends it to a configurable
collections client.

## Supported statistics

- Memory usage
- Goroutines
- Open files
- Garbage Collection
- CGO Calls

## Supported collector clients

- statsd

## Example

```go
func main() {
	statsd := statsd.New()
	stats := stats.New(stats.Client(statsd))
	go stats.Collect()
}
```
