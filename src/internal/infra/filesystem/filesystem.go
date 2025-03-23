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

		fileExtension := filepath.Ext(entry.Name())
		fileName := entry.Name()[:len(entry.Name())-len(fileExtension)]

		if entry.IsDir() {
			continue
		}

		if extension != "" && extension != fileExtension {
			continue
		}

		if entry.Name()[0] == '.' {
			continue
		}

		files = append(files, *entity.NewFile(fileName, fileExtension, folder))
	}

	return files, nil

}
