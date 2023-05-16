# configstore

[![Go Reference](https://pkg.go.dev/badge/github.com/vmorsell/configstore.svg)](https://pkg.go.dev/github.com/vmorsell/configstore)

Persistent config storage for Go applications. Allows putting and getting a single config
file.

## Installation

Install configstore using `go get`

    go get github.com/vmorsell/configstore

Import the package in your app

```go
import (
    "github.com/vmorsell/configstore"
)

func main() {
    store := configstore.Must(configstore.New("my-app"))
    /* ... */
}
```

## Example usage

```go
type Config struct {
    APIKey string `json:"api_key"`
}

store := configstore.Must(configstore.New("my-app"))

// Write config to file.
config := Config{
    APIKey: "xyz",
}
if err := store.Put(config); err != nil {
    panic(err)
}

// Read config from file.
stored := Config{}
if err := store.Get(&stored); err != nil {
    panic(err)
}

fmt.Println(stored.APIKey)
```
