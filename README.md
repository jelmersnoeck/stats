# Stats

[![Build Status](https://travis-ci.org/jelmersnoeck/stats.svg?branch=master)](https://travis-ci.org/jelmersnoeck/stats)
[![GoDoc](https://godoc.org/github.com/jelmersnoeck/stats?status.svg)](https://godoc.org/github.com/jelmersnoeck/stats)

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
package main

import (
    "github.com/jelmersnoeck/stats"
    "github.com/jelmersnoeck/stats/clients/statsd"

    sd "gopkg.in/alexcesaro/statsd.v2"
)

func main() {
    cl, err := sd.New()
    if err != nil {
        // error connecting to statsd
    }

    stats := stats.New(stats.Client(statsd.New(cl)))
    go stats.Collect()
}
```
