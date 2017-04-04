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
	os.Mkdir(filepath.Join(path, "subDir1"), 777)
	os.Mkdir(filepath.Join(path, "subDir2"), 777)

	dir, err := ScanPath(path)

	require.NoError(t, err)
	require.Equal(t, Dir{
		Path:  path,
		Files: []File{{"1.txt"}, {"2.mp3"}},
		Dirs: []Dir{
			{Path: filepath.Join(path, "subDir1")},
			{Path: filepath.Join(path, "subDir2")},
		},
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
	require.Empty(t, dir.Files)
	require.Empty(t, dir.Dirs)
}

func TestScanPath_givenNonExistingDir(t *testing.T) {
	path := "unknown-dir"

	_, err := ScanPath(path)

	require.Error(t, err)
}

func TestDir_ContainsMusic_givenMp3(t *testing.T) {
	folder := Dir{"album", []File{{"foo.mp3"}}, []Dir{}}

	require.True(t, folder.ContainsMusic())
}

func TestDir_ContainsMusic_givenNoMusicFile(t *testing.T) {
	folder := Dir{"album", []File{{"readme.txt"}}, []Dir{}}

	require.False(t, folder.ContainsMusic())
}

func TestDir_ContainsMusic_givenSubDirectory(t *testing.T) {
	folder := Dir{"album", []File{{"readme.txt"}}, []Dir{{Path: "sub"}}}

	require.False(t, folder.ContainsMusic())
}

func TestDir_ContainsMusic_givenEmptyDir(t *testing.T) {
	folder := Dir{"album", []File{}, []Dir{}}

	require.False(t, folder.ContainsMusic())
}

func TestDir_ContainsPlaylist_givenM3u(t *testing.T) {
	folder := Dir{"album", []File{{"a playlist.m3u"}}, []Dir{}}

	contains, playlist := folder.ContainsPlaylist()

	require.True(t, contains)
	require.Equal(t, File{"a playlist.m3u"}, playlist)
}

func TestDir_ContainsPlaylist_givenMusicButNoPlaylist(t *testing.T) {
	folder := Dir{"album", []File{{"foo.mp3"}, {"bar.mp3"}}, []Dir{}}

	contains, playlist := folder.ContainsPlaylist()

	require.False(t, contains)
	require.Zero(t, playlist)
}

func TestDir_ContainsPlaylist_givenSubDirectory(t *testing.T) {
	folder := Dir{"album", []File{}, []Dir{{Path: "sub"}}}

	contains, playlist := folder.ContainsPlaylist()

	require.False(t, contains)
	require.Zero(t, playlist)
}

func TestDir_ContainsPlaylist_givenEmptyDir(t *testing.T) {
	folder := Dir{"album", []File{}, []Dir{}}

	contains, playlist := folder.ContainsPlaylist()

	require.False(t, contains)
	require.Zero(t, playlist)
}
