package dir

import (
	"io/ioutil"
)

type Dir struct {
	Path  string
	Files []File
}

func ScanPath(path string) (Dir, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return Dir{}, err
	}

	dir := Dir{
		Path: path,
	}

	for _, info := range fileInfos {
		if !info.IsDir() {
			dir.Files = append(dir.Files, File{info.Name()})
		}
	}

	return dir, nil
}

func (d Dir) ContainsMusic() bool {
	for _, file := range d.Files {
		if file.IsMusic() {
			return true
		}
	}
	return false
}

func (d Dir) ContainsPlaylist() (bool, File) {
	for _, file := range d.Files {
		if file.IsPlaylist() {
			return true, file
		}
	}
	return false, File{}
}
