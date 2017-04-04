package playlist

import (
	"github.com/stretchr/testify/require"
	"github.com/tomsquest/goplaylistcreator/dir"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestPlaylist_Create(t *testing.T) {
	folder := createDirWithFiles(t, "1.mp3", "2.mp3")
	defer os.RemoveAll(folder.Path())

	result, err := Create(folder)

	require.NoError(t, err)
	require.Equal(t, Created, result.Status)
	requireContent(t, "1.mp3\n2.mp3\n", result)
}

func TestCreate_ignoresNonMusicFiles(t *testing.T) {
	folder := createDirWithFiles(t, "readme.txt", "1.mp3")
	defer os.RemoveAll(folder.Path())

	result, err := Create(folder)

	require.NoError(t, err)
	require.Equal(t, Created, result.Status)
	requireContent(t, "1.mp3\n", result)
}

func TestCreate_doesNotCreate_givenExistingPlaylist(t *testing.T) {
	folder := createDirWithFiles(t, "aPlaylist.pls", "1.mp3")
	defer os.RemoveAll(folder.Path())

	result, err := Create(folder)

	require.NoError(t, err)
	require.Equal(t, AlreadyExisting, result.Status)
	require.Equal(t, filepath.Join(folder.Path(), "aPlaylist.pls"), result.PlaylistPath)
}

func createDirWithFiles(t *testing.T, filenames ...string) dir.Dir {
	path, err := ioutil.TempDir("", "test-playlistcreator")
	if err != nil {
		t.Fatal(err)
	}

	for _, filename := range filenames {
		err := ioutil.WriteFile(filepath.Join(path, filename), []byte{}, 0666)
		if err != nil {
			t.Fatal(err)
		}
	}

	folder, err := dir.ScanPath(path)
	if err != nil {
		t.Fatal(err)
	}

	return folder
}

func requireContent(t *testing.T, expectedContent string, actualResult CreationResult) {
	actualContent, err := ioutil.ReadFile(actualResult.PlaylistPath)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, expectedContent, string(actualContent))
}
