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
	rootPath := flag.Arg(0)

	if rootPath == "" {
		log.Fatal("Missing root folder as first argument")
	}

	root, err := dir.ScanPath(rootPath)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Fatalf("Folder not found: \"%s\"\n", rootPath)
		} else {
			log.Fatalf("Unable to scan folder: \"%s\". Reason: %s", rootPath, err)
		}
	}

	result, err := playlist.Create(root)
	if err != nil {
		log.Fatal(err)
	}
	if result.Status == playlist.Created {
		log.Printf("Playlist created: %s\n", result.PlaylistPath)
	} else if result.Status == playlist.AlreadyExisting {
		log.Printf("Already existing playlist: %s\n", result.PlaylistPath)
	}
}
