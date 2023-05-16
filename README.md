# config-store

[![Go Reference](https://pkg.go.dev/badge/github.com/vmorsell/config-store.svg)](https://pkg.go.dev/github.com/vmorsell/config-store)

Persistent config storage with go. Allows putting and getting a single config
file for an application.

## Installation

Install config-store using `go get`

    go get github.com/vmorsell/config-store

Import the package in your app

```go
import (
    "github.com/vmorsell/config-store"
)

func main() {
    configStore := MustNewConfigStore().WithAppName("my-app")
    /* ... */
}
```

## Example usage

```go
type Config struct {
    APIKey string `json:"apiKey"`
    Username string `json:"apiKey"`
}

configStore := configstore.MustNewConfigStore().WithAppName("my-app")

config := Config{
    APIKey: "abc",
    Username: "user1",
}

store.Put(config)

stored := Config{}
store.Get(&stored)
```
