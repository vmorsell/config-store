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

func TestNewConfigStore(t *testing.T) {
	_, err := New("app")
	require.Nil(t, err)
}

func TestMustNewConfigStore(t *testing.T) {
	require.NotPanics(t, func() {
		Must(New("app"))
	})
}

func TestDir(t *testing.T) {
	appName := "test2"
	rootDir := os.TempDir()
	want := filepath.Join(rootDir, appName)

	store, err := New(appName)
	require.Nil(t, err)

	store.RootDir = rootDir

	dir := store.Dir()
	require.Equal(t, dir, want)
}

func TestFilepath(t *testing.T) {
	appName := "test3"
	rootDir := os.TempDir()
	want := filepath.Join(rootDir, appName, ConfigFilename)

	store, err := New(appName)
	require.Nil(t, err)

	store.RootDir = rootDir

	path := store.Filepath()
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
