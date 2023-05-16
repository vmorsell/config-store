package configstore

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

type Config struct {
	A string  `json:"a"`
	B *string `json:"b"`
}

func TestGetPut(t *testing.T) {
	appName := "test1"
	store := MustNewConfigStore().
		WithAppName(appName).
		WithRootDir(t.TempDir())

	config := Config{
		A: "hello",
	}

	store.Put(config)

	res := Config{}
	store.Get(&res)

	require.Equal(t, res, config)
}

func TestDir(t *testing.T) {
	rootDir := os.TempDir()
	appName := "test2"
	want := filepath.Join(rootDir, configDir, appName)

	S := MustNewConfigStore().
		WithAppName(appName).
		WithRootDir(rootDir)

	s := S.(*configStore)

	dir := s.dir()
	require.Equal(t, dir, want)
}

func TestFilepath(t *testing.T) {
	rootDir := os.TempDir()
	appName := "test3"
	want := filepath.Join(rootDir, configDir, appName, configFilename)

	S := MustNewConfigStore().
		WithRootDir(rootDir).
		WithAppName(appName)

	s := S.(*configStore)

	path := s.filepath()
	require.Equal(t, want, path)
}

func TestEnsureDirExists(t *testing.T) {
	a := "a"
	path := filepath.Join(t.TempDir(), a)

	_, err := os.Stat(path)
	require.ErrorIs(t, err, os.ErrNotExist)

	err = ensureDirExists(path)
	require.Nil(t, err)

	_, err = os.Stat(path)
	require.Nil(t, err)
}
