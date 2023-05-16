// Package configstore provides a persistent storage for config files. It
// supports storing config to a single JSON file per application.
package configstore

import (
	"encoding/json"
	"fmt"
	"io"
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

// NewConfigStore returns a new config store.
func New(appName string) (*ConfigStore, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("user config dir: %w", err)
	}

	if appName == "" {
		appName = DefaultAppName
	}

	return &ConfigStore{
		AppName: appName,
		RootDir: dir,
	}, nil
}

// Must is a helper function to ensure the config store is valid and there was
// no error when calling a NewConfigStore function.
//
// This helper is intended to be used in variable initialization to load the
// Session and configuration at startup. Such as:
//
//	store := configstore.Must(configstore.New("app_name"))
func Must(cs *ConfigStore, err error) *ConfigStore {
	if err != nil {
		panic(err)
	}

	return cs
}

// Get reads the config file from disk and stores the config in the value
// pointed to by v.
func (cs *ConfigStore) Get(v interface{}) error {
	if err := ensureDirExists(cs.Dir()); err != nil {
		return fmt.Errorf("ensure dir exists: %w", err)
	}

	file, err := os.Open(cs.Filepath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("open file: %w", err)
	}

	data, err := io.ReadAll(file)
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
func (cs *ConfigStore) Put(v interface{}) error {
	if err := ensureDirExists(cs.Dir()); err != nil {
		return fmt.Errorf("ensure dir exists: %w", err)
	}

	file, err := os.Create(cs.Filepath())
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
