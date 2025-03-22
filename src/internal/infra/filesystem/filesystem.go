package filesystem

import (
	"os"
	"path/filepath"

	"github.com/luizhp/query-extract/internal/entity"
)

func ListFolder(folder string, extension string) ([]entity.File, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	var files []entity.File
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

		files = append(files, entity.File{
			Name: entry.Name(),
			Path: filepath.Join(folder, entry.Name()),
		})
	}

	return files, nil

}
