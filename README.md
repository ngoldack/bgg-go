# bgg-go

![CI](https://github.com/ngoldack/bgg-go/actions/workflows/ci.yaml/badge.svg)
![Security Check](https://github.com/ngoldack/bgg-go/actions/workflows/codeql.yaml/badge.svg)
[![codecov](https://codecov.io/github/ngoldack/bgg-go/branch/main/graph/badge.svg?token=59RN2REWTT)](https://codecov.io/github/ngoldack/bgg-go)
[![Go Report Card](https://goreportcard.com/badge/ngoldack/bgg-go)](https://goreportcard.com/report/github.com/ngoldack/bgg-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/ngoldack/bgg-go.svg)](https://pkg.go.dev/github.com/ngoldack/bgg-go)

unofficial go library for [boardgamegeek](https://boardgamegeek.com)

**IMPORTANT: This library is still in active development, do not consider it as stable! Documentation may be incomplete!**

## Installation

```shell
go get -u github.com/ngoldack/bgg-go
```

## Usage

```go
package main

import "github.com/ngoldack/bgg-go/bgg"

func main() {
    // Create a new client
    client := bgg.New()

    // Search for games with the name 'Chess' 
    results, err := client.Search("Chess", nil)
    if err != nil {
        panic(err)
    }

    // do something with the results
}
```

## Roadmap

- [ ] basic XML-API2 feature implementation
  - [X] search
  - [X] thing
  - [ ] user
  - [ ] collection
  - [ ] guilds
  - [ ] plays
  - [ ] hot
  - [ ] forums
  - [ ] threads
- [ ] combined helper functions for easy searching with richer results

## Contributing

See the [Contributing.md](Contributing.md)