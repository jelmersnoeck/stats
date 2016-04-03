# Stats

Stats is a statistics collection library for your Go applications. It collects
data on a regular, configurable interval and sends it to a configurable
collections client.

## Supported statistics

- Memory usage
- Goroutines
- Open files
- Garbage Collection
- CGO Calls

## Supported collectors

The collector interface has only one method, `Count(bucket string, count int64)`.
