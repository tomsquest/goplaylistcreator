package dir

import (
	"path/filepath"
	"strings"
)

type File struct {
	Name string
}

func (file File) IsMusic() bool {
	switch strings.ToLower(filepath.Ext(file.Name)) {
	case ".mp3", ".mp4a", ".ogg":
		return true
	}
	return false
}

func (file File) IsPlaylist() bool {
	switch strings.ToLower(filepath.Ext(file.Name)) {
	case ".m3u", ".m3u8", ".pls":
		return true
	}
	return false
}
