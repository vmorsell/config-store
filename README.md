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
    configStore := MustNewConfigStore("my-app")
    /* ... */
}
```

## Example usage

```go
type Config struct {
    APIKey string `json:"api_key"`
    Username string `json:"username"`
}

configStore := configstore.MustNewConfigStore("my-app")

config := Config{
    APIKey: "abc",
    Username: "user1",
}

store.Put(config)

stored := Config{}
store.Get(&stored)

fmt.Println(stored.APIKey)
```
