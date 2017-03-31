package main

import (
	"flag"
	"github.com/tomsquest/goplaylistcreator/dir"
	"github.com/tomsquest/goplaylistcreator/playlist"
	"log"
	"os"
)

func main() {
	flag.Parse()
	folderPath := flag.Arg(0)

	if folderPath == "" {
		log.Fatal("Missing target folder as first argument")
	}

	folder, err := dir.ScanPath(folderPath)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Fatalf("Folder not found: \"%s\"\n", folderPath)
		} else {
			log.Fatalf("Unable to scan folder: \"%s\". Reason: %s", folderPath, err)
		}
	}

	result, err := playlist.Create(folder)
	if err != nil {
		log.Fatal(err)
	}
	if result.Status == playlist.Created {
		log.Printf("Playlist created: %s\n", result.PlaylistPath)
	} else if result.Status == playlist.AlreadyExisting {
		log.Printf("Already existing playlist: %s\n", result.PlaylistPath)
	}
}
