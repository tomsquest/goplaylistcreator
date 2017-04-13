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
	levelB1 := filepath.Join(rootPath, "levelA", "levelB1")
	levelB1C := filepath.Join(rootPath, "levelA", "levelB1", "levelC")
	os.MkdirAll(levelB1C, 0770)
	ioutil.WriteFile(filepath.Join(levelA, "levelA.mp3"), []byte{}, 0666)
	ioutil.WriteFile(filepath.Join(levelB1, "levelB.mp3"), []byte{}, 0666)
	ioutil.WriteFile(filepath.Join(levelB1C, "levelC.mp3"), []byte{}, 0666)
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

	requirePlaylistExist(t, levelA, "levelA.m3u")
	requirePlaylistExist(t, levelB1, "levelB1.m3u")
	requirePlaylistExist(t, levelB1C, "levelC.m3u")
	requireNoPlaylist(t, levelB2)
	requireNoPlaylist(t, levelC2)
}

func requirePlaylistExist(t *testing.T, playlistDir string, expectedPlaylistName string) {
	info, err := os.Stat(filepath.Join(playlistDir, expectedPlaylistName))
	require.NoError(t, err)
	require.False(t, info.IsDir())
}

func requireNoPlaylist(t *testing.T, path string) {
	files, err := ioutil.ReadDir(path)
	require.NoError(t, err)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".m3u" {
			require.Fail(t, "Unexpected playlist found", file.Name())
		}
	}
	//_, err := os.Stat(filepath.Join(path, "playlist.m3u"))
}
