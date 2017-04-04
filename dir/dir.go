package dir

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type Dir struct {
	path          string
	files         []File
	dirs          []Dir
	containsMusic bool
	playlist      playlist
}

type playlist struct {
	Exist bool
	Path  string
}

func ScanPath(path string) (Dir, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return Dir{}, err
	}

	dir := Dir{
		path: path,
	}

	for _, info := range fileInfos {
		subPath := filepath.Join(path, info.Name())
		if info.IsDir() {
			subDir, err := ScanPath(subPath)
			if err != nil {
				log.Printf("Skipping folder: %s. Reason: %s", subPath, err)
			} else {
				dir.dirs = append(dir.dirs, subDir)
			}
		} else {
			file := File{info.Name()}
			dir.files = append(dir.files, file)

			if file.IsMusic() {
				dir.containsMusic = true
			}
			if file.IsPlaylist() {
				dir.playlist = playlist{true, subPath}
			}
		}
	}

	return dir, nil
}

func (d *Dir) Path() string {
	return d.path
}

func (d *Dir) Files() []File {
	return d.files
}

func (d *Dir) Dirs() []Dir {
	return d.dirs
}

func (d *Dir) ContainsMusic() bool {
	return d.containsMusic
}

func (d *Dir) Playlist() playlist {
	return d.playlist
}
