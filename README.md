# WaitGroup - Parallel-Controlled WaitGroup

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/waitgroup)](https://pkg.go.dev/github.com/go-zoox/waitgroup)
[![Build Status](https://github.com/go-zoox/waitgroup/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/waitgroup/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/waitgroup)](https://goreportcard.com/report/github.com/go-zoox/waitgroup)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/waitgroup/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/waitgroup?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/waitgroup.svg)](https://github.com/go-zoox/waitgroup/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/waitgroup.svg?label=Release)](https://github.com/go-zoox/waitgroup/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/waitgroup
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/waitgroup"
)

func main(t *testing.T) {
  wg := waitgroup.New(3)
  jobs := []func(){}

  for i := 0; i < 10; i++ {
    index := i
    jobs = append(jobs, func() {
      time.Sleep(time.Second)
      fmt.Println(index)
    })
  }

  wg.Add(jobs...)
  
  wg.Wait()
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
