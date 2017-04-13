package playlist

import (
	"github.com/pkg/errors"
	"github.com/tomsquest/goplaylistcreator/dir"
	"os"
	"path/filepath"
)

type Playlist struct{}

type CreationResult struct {
	Status       code
	PlaylistPath string
}

type code int

const (
	Created code = iota
	AlreadyExisting
	NoMusicFiles
)

func Create(dir dir.Dir) (CreationResult, error) {
	previousPlaylist := dir.Playlist()
	if previousPlaylist.Exist {
		return CreationResult{
			Status:       AlreadyExisting,
			PlaylistPath: previousPlaylist.Path,
		}, nil
	}

	if !dir.ContainsMusic() {
		return CreationResult{
			Status: NoMusicFiles,
		}, nil
	}

	path := filepath.Join(dir.Path(), filepath.Base(dir.Path())+".m3u")
	playlist, err := os.Create(path)
	if err != nil {
		return CreationResult{}, errors.Errorf("Unable to create playlist: %s. Reason: %s\n", path, err)
	}
	defer playlist.Close()

	for _, file := range dir.Files() {
		if file.IsMusic() {
			playlist.WriteString(file.Name + "\n")
		}
	}

	result := CreationResult{
		Status:       Created,
		PlaylistPath: path,
	}

	return result, nil
}
