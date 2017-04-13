package playlist

import (
	"github.com/tomsquest/goplaylistcreator/dir"
	"log"
)

func CreateAll(root dir.Dir) {
	result, err := Create(root)
	if err != nil {
		log.Printf("   [Error] %s", err)
	} else {
		if result.Status == Created {
			log.Printf(" [Success] Playlist created: %s", result.PlaylistPath)
		} else if result.Status == AlreadyExisting {
			log.Printf("[Skipping] Already existing playlist: %s", result.PlaylistPath)
		} else if result.Status == NoMusicFiles {
			log.Printf("[Skipping] No music files in %s", root.Path())
		}
	}

	for _, subDir := range root.Dirs() {
		CreateAll(subDir)
	}
}
