package dir

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestScanPath_givenSomeFilesAndDirAtRoot(t *testing.T) {
	path, _ := ioutil.TempDir("", "test-playlistcreator")
	defer os.RemoveAll(path)
	ioutil.WriteFile(filepath.Join(path, "1.txt"), nil, 0666)
	ioutil.WriteFile(filepath.Join(path, "2.mp3"), nil, 0666)
	os.Mkdir(filepath.Join(path, "subDir1"), 0777)
	os.Mkdir(filepath.Join(path, "subDir2"), 0777)

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Equal(t, Dir{
		path:  path,
		files: []File{{"1.txt"}, {"2.mp3"}},
		dirs: []Dir{
			{path: filepath.Join(path, "subDir1")},
			{path: filepath.Join(path, "subDir2")},
		},
		containsMusic: true,
	}, dir)
}

func TestScanPath_scansRecursively(t *testing.T) {
	path, _ := ioutil.TempDir("", "test-playlistcreator")
	defer os.RemoveAll(path)
	subDir := filepath.Join(path, "subDir")
	os.Mkdir(subDir, 0777)
	subSubDir := filepath.Join(subDir, "subSubDir")
	os.Mkdir(subSubDir, 0777)
	ioutil.WriteFile(filepath.Join(path, "file.mp3"), nil, 0666)
	ioutil.WriteFile(filepath.Join(subDir, "subFile.mp3"), nil, 0666)
	ioutil.WriteFile(filepath.Join(subSubDir, "subSubFile.mp3"), nil, 0666)

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Equal(t, Dir{
		path:  path,
		files: []File{{"file.mp3"}},
		dirs: []Dir{
			{
				path: subDir,
				dirs: []Dir{
					{
						path:          subSubDir,
						files:         []File{{"subSubFile.mp3"}},
						containsMusic: true,
					},
				},
				files:         []File{{"subFile.mp3"}},
				containsMusic: true,
			},
		},
		containsMusic: true,
	}, dir)
}

func TestScanPath_givenEmptyDir(t *testing.T) {
	path, err := ioutil.TempDir("", "test-playlistcreator")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(path)

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Empty(t, dir.Files())
	require.Empty(t, dir.Dirs())
	require.False(t, dir.ContainsMusic())
}

func TestScanPath_givenNonExistingDir(t *testing.T) {
	path := "unknown-dir"

	_, err := ScanPath(path)

	require.Error(t, err)
}

func TestDir_ContainsMusic_givenMp3(t *testing.T) {
	path, _ := ioutil.TempDir("", "test-playlistcreator")
	defer os.RemoveAll(path)
	ioutil.WriteFile(filepath.Join(path, "song.mp3"), nil, 0666)

	dir, _ := ScanPath(path)

	require.True(t, dir.ContainsMusic())
}

func TestDir_ContainsMusic_givenNoMusicFile(t *testing.T) {
	path, _ := ioutil.TempDir("", "test-playlistcreator")
	defer os.RemoveAll(path)
	ioutil.WriteFile(filepath.Join(path, "readme.md"), nil, 0666)
	os.Mkdir(filepath.Join(path, "subDir"), 0777)

	dir, _ := ScanPath(path)

	require.False(t, dir.ContainsMusic())
}

func TestDir_ContainsPlaylist_givenAPlaylist(t *testing.T) {
	path, _ := ioutil.TempDir("", "test-playlistcreator")
	defer os.RemoveAll(path)
	playlist := filepath.Join(path, "playlist.m3u")
	ioutil.WriteFile(playlist, nil, 0666)

	dir, _ := ScanPath(path)

	require.True(t, dir.Playlist().Exist)
	require.Equal(t, playlist, dir.Playlist().Path)
}

func TestDir_ContainsPlaylist_givenNoPlaylist(t *testing.T) {
	path, _ := ioutil.TempDir("", "test-playlistcreator")
	defer os.RemoveAll(path)
	ioutil.WriteFile(filepath.Join(path, "song.mp3"), nil, 0666)
	os.Mkdir(filepath.Join(path, "subDir"), 0777)

	dir, _ := ScanPath(path)

	require.False(t, dir.Playlist().Exist)
	require.Empty(t, dir.Playlist().Path)
}
