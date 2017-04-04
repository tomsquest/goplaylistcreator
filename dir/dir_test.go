package dir

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestScanPath_givenSomeFiles(t *testing.T) {
	path, err := ioutil.TempDir("", "test-playlistcreator")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(path)
	ioutil.WriteFile(filepath.Join(path, "1.txt"), nil, 0666)
	ioutil.WriteFile(filepath.Join(path, "2.mp3"), nil, 0666)

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Equal(t, []File{{"1.txt"}, {"2.mp3"}}, dir.Files)
}

func TestScanPath_givenEmptyDir(t *testing.T) {
	path, err := ioutil.TempDir("", "test-playlistcreator")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(path)

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Empty(t, dir.Files)
}

func TestScanPath_givenNonExistingDir(t *testing.T) {
	path := "unknown-dir"

	_, err := ScanPath(path)

	require.Error(t, err)
}

func TestScanPath_givenSubDirectory(t *testing.T) {
	path, err := ioutil.TempDir("", "test-playlistcreator")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(path)
	ioutil.TempDir(path, "subdir")

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Empty(t, dir.Files)
}

func TestDir_ContainsMusic_givenMp3(t *testing.T) {
	folder := Dir{"album", []File{{"foo.mp3"}}}

	require.True(t, folder.ContainsMusic())
}

func TestDir_ContainsMusic_givenNoMusicFile(t *testing.T) {
	folder := Dir{"album", []File{{"readme.txt"}}}

	require.False(t, folder.ContainsMusic())
}

func TestDir_ContainsMusic_givenEmptyDir(t *testing.T) {
	folder := Dir{"album", []File{}}

	require.False(t, folder.ContainsMusic())
}

func TestDir_ContainsPlaylist_givenM3u(t *testing.T) {
	folder := Dir{"album", []File{{"a playlist.m3u"}}}

	contains, file := folder.ContainsPlaylist()

	require.True(t, contains)
	require.Equal(t, File{"a playlist.m3u"}, file)
}

func TestDir_ContainsPlaylist_givenMusicButNoPlaylist(t *testing.T) {
	folder := Dir{"album", []File{{"foo.mp3"}, {"bar.mp3"}}}

	contains, file := folder.ContainsPlaylist()

	require.False(t, contains)
	require.Zero(t, file)
}

func TestDir_ContainsPlaylist_givenEmptyDir(t *testing.T) {
	folder := Dir{"album", []File{}}

	contains, file := folder.ContainsPlaylist()

	require.False(t, contains)
	require.Zero(t, file)
}
