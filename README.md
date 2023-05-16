# config-store

[![Go Reference](https://pkg.go.dev/badge/github.com/vmorsell/config-store.svg)](https://pkg.go.dev/github.com/vmorsell/config-store)

Persistent config storage for Go applications. Allows putting and getting a single config
file.

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
    APIKey string `json:"api_key"`
    Username string `json:"username"`
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
