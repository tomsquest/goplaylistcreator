package playlist

import (
	"github.com/stretchr/testify/require"
	"github.com/tomsquest/goplaylistcreator/dir"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateInAllDirs(t *testing.T) {
	rootPath, _ := ioutil.TempDir("", "createmany")
	//defer os.RemoveAll(rootPath)

	// Directories with music
	levelA := filepath.Join(rootPath, "levelA")
	levelB := filepath.Join(rootPath, "levelA", "levelB1")
	levelC := filepath.Join(rootPath, "levelA", "levelB1", "levelC")
	os.MkdirAll(levelC, 0770)
	ioutil.WriteFile(filepath.Join(levelA, "levelA.mp3"), []byte{}, 0666)
	ioutil.WriteFile(filepath.Join(levelB, "levelB.mp3"), []byte{}, 0666)
	ioutil.WriteFile(filepath.Join(levelC, "levelC.mp3"), []byte{}, 0666)
	// Empty directories
	levelB2 := filepath.Join(rootPath, "levelA", "levelB2")
	levelC2 := filepath.Join(rootPath, "levelA", "levelB2", "levelC2")
	os.MkdirAll(levelB2, 0770)
	os.MkdirAll(levelC2, 0770)

	root, err := dir.ScanPath(rootPath)
	if err != nil {
		t.Fatal(err)
	}

	CreateAll(root)

	requirePlaylistExist(t, levelA)
	requirePlaylistExist(t, levelB)
	requirePlaylistExist(t, levelC)
	requireNoPlaylist(t, levelB2)
	requireNoPlaylist(t, levelC2)
}

func requirePlaylistExist(t *testing.T, path string) {
	info, err := os.Stat(filepath.Join(path, "playlist.m3u"))
	require.NoError(t, err)
	require.False(t, info.IsDir())
}

func requireNoPlaylist(t *testing.T, path string) {
	_, err := os.Stat(filepath.Join(path, "playlist.m3u"))
	require.Error(t, err)
}
