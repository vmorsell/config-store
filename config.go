package configstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ConfigStore is the interface for the config.
type ConfigStore interface {
	WithAppName(name string) ConfigStore
	WithRootDir(dir string) ConfigStore

	Get(v interface{}) error
	Put(config interface{}) error
}

type configStore struct {
	rootDir string
	appName string
}

const (
	configDir      = ".config"
	configFilename = "config.json"
	defaultAppName = "default"
)

type Settings struct {
	appName string
}

// NewConfigStore returns a new config.
func NewConfigStore() (ConfigStore, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("user home dir: %w", err)
	}

	return &configStore{
		rootDir: filepath.Join(homeDir, configDir),
		appName: defaultAppName,
	}, nil
}

// MustNewConfig returns a new config or panics.
func MustNewConfigStore() ConfigStore {
	configStore, err := NewConfigStore()
	if err != nil {
		panic(err)
	}
	return configStore
}

// Set app name. This controls where the config file is stored.
func (c *configStore) WithAppName(name string) ConfigStore {
	c.appName = name
	return c
}

// Set set root dir. This defaults to $HOME/.config on Unix and macOS,
// and %USERPROFILE%/.config on Windows.
// You probably don't want to do this.
func (c *configStore) WithRootDir(dir string) ConfigStore {
	c.rootDir = filepath.Join(dir, configDir)
	return c
}

var (
	errAppNameNotSet = fmt.Errorf("app name not set")
)

// Get a config and store the result in the value pointed to by v.
func (c *configStore) Get(v interface{}) error {
	if err := ensureDirExists(c.dir()); err != nil {
		return fmt.Errorf("ensure dir exists: %w", err)
	}

	file, err := os.Open(c.filepath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("open file: %w", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

// Put the provided config.
func (c *configStore) Put(v interface{}) error {
	if err := ensureDirExists(c.dir()); err != nil {
		return fmt.Errorf("ensure dir exists: %w", err)
	}

	file, err := os.Create(c.filepath())
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

// dir returns the full path to the directory where the config file is stored.
func (c *configStore) dir() string {
	return filepath.Join(c.rootDir, c.appName)
}

// filepath returns the full path to the config file.
func (c *configStore) filepath() string {
	return filepath.Join(c.dir(), configFilename)
}

// ensureDirExists ensures all folders in the path exists.
func ensureDirExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat: %w", err)
		}
		os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
