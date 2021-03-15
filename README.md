# readenv

A simple Go library to read environment variables files(.env files).

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
