package configstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigStore struct {
	// The app name controls the name of the directory where the config JSON
	// file will be stored. The path will look something like the path below.
	// root_dir/.config/app_name/config.json
	AppName string

	// RootDir defines where config files are stored in the operating system.
	// This defaults to $HOME/.config on Unix and macOS, and
	// %USERPROFILE%/.config on Windows.
	// You probably want to go with the default from NewConfigStore().
	RootDir string
}

const (
	// Configfilename is the name of the JSON file where the config is stored.
	ConfigFilename = "config.json"

	// DefaultAppName is the fallback name for the application directory
	// unless set.
	DefaultAppName = "unnamed_app"
)

// NewConfigStore returns a new config.
func NewConfigStore(appName string) (*ConfigStore, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("user config dir: %w", err)
	}

	if appName == "" {
		appName = defaultAppName
	}

	return &ConfigStore{
		AppName: appName,
		RootDir: dir,
	}, nil
}

// MustNewConfig returns a new config or panics.
func MustNewConfigStore(appName string) *ConfigStore {
	configStore, err := NewConfigStore(appName)
	if err != nil {
		panic(err)
	}
	return configStore
}

var (
	errAppNameNotSet = fmt.Errorf("app name not set")
)

// Get reads the config file from disk and stores the config in the value
// pointed to by v.
func (c *ConfigStore) Get(v interface{}) error {
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

// Put writes the provided configuration to disk. If a config file already
// exists, it will be overwritten.
func (c *ConfigStore) Put(v interface{}) error {
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

// Dir returns the full path to the directory where the config file is stored.
func (cs *ConfigStore) Dir() string {
	return filepath.Join(cs.RootDir, cs.AppName)
}

// Filepath returns the full path to the config file.
func (cs *ConfigStore) Filepath() string {
	return filepath.Join(cs.Dir(), ConfigFilename)
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
