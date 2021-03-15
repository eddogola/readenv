# readenv

[![codecov](https://codecov.io/gh/eddogola/readenv/branch/main/graph/badge.svg?token=SLKBLF3ZDW)](https://codecov.io/gh/eddogola/readenv)
[![Go Reference](https://pkg.go.dev/badge/github.com/eddogola/readenv.svg)](https://pkg.go.dev/github.com/eddogola/readenv)
[![Build Status](https://travis-ci.com/eddogola/readenv.svg?branch=main)](https://travis-ci.com/eddogola/readenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/eddogola/readenv)](https://goreportcard.com/report/github.com/eddogola/readenv)

A simple Go library to read environment variables files(.env files).

## Installing

```bash
go get -u github.com/eddogola/readenv
```

## Usage

The library basically parses the (file) byte data into a map, from which users can Get
environment variables.

```go
import (
    "os"

    "github.com/eddogola/readenv"
)

func main() {
    file, err := os.Open(".env", os.O_RDONLY, 0444) // open .env file
    if err != nil {
        // handle error
    }
    defer file.Close()
    
    envData, err := readenv.ReadEnv(file)
    if err != nil {
        // handle error
    }

    val, err := envData.Get("<YOUR_.ENV_VAR>")) // access specific variables using their keys
    if err != nil {
        // handle error
    }
    fmt.Println(val)
}
```

## Contribution

Feel free to

- Start issues.
- Make pull requests.
- Make suggestions on improvement, e.g. features
