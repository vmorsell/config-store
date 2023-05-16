package configstore_test

import (
	"fmt"

	"github.com/vmorsell/configstore"
)

type Config struct {
	APIKey string `json:"api_key"`
}

func Example() {
	store := configstore.Must(configstore.New("app"))

	// Store config to file.
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
}
