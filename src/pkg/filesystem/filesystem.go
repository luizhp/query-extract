package filesystem

import (
	"os"
	"path/filepath"
)

type File struct {
	Name string
	Path string
}

func ListFolder(folder string, extension string) ([]File, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	var files []File
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if extension != "" && extension != filepath.Ext(entry.Name()) {
			continue
		}

		if entry.Name()[0] == '.' {
			continue
		}

		files = append(files, File{
			Name: entry.Name(),
			Path: filepath.Join(folder, entry.Name()),
		})
	}

	return files, nil

}
