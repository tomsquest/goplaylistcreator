package dir

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type Dir struct {
	path  string
	files []File
	dirs  []Dir
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
		if info.IsDir() {
			subPath := filepath.Join(path, info.Name())
			subDir, err := ScanPath(subPath)
			if err != nil {
				log.Printf("Skipping folder: %s. Reason: %s", subPath, err)
			} else {
				dir.dirs = append(dir.dirs, subDir)
			}
		} else {
			dir.files = append(dir.files, File{info.Name()})
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
	for _, file := range d.Files() {
		if file.IsMusic() {
			return true
		}
	}
	return false
}

func (d *Dir) ContainsPlaylist() (bool, File) {
	for _, file := range d.Files() {
		if file.IsPlaylist() {
			return true, file
		}
	}
	return false, File{}
}
