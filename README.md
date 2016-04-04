# Stats

[![Build Status](https://travis-ci.org/jelmersnoeck/stats.svg?branch=master)](https://travis-ci.org/jelmersnoeck/stats)
[![GoDoc](https://godoc.org/github.com/jelmersnoeck/stats?status.svg)](https://godoc.org/github.com/jelmersnoeck/stats)

Stats is a statistics collection library for your Go applications. It collects
data on a regular, configurable interval and sends it to a configurable
collections client.

## Supported statistics

- Memory usage
- Goroutines
- Garbage Collection
- CGO Calls

## Supported collector clients

- statsd

## Example

```go
package main

import (
    "github.com/jelmersnoeck/stats"
    "github.com/jelmersnoeck/stats/clients/statsd"
)

func main() {
    cl, err := statsd.New()
    if err != nil {
        // error connecting to statsd
    }

    s := stats.New(stats.Client(cl))
    go s.Collect()
}
```
