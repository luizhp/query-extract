package filesystem

import (
	"io"
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

		fileName := entry.Name()
		fileExtension := filepath.Ext(entry.Name())
		fileName = fileName[:len(fileName)-len(fileExtension)]

		if fileExtension != "" {
			fileExtension = fileExtension[1:]
		}
		// fileName := entry.Name()[:len(entry.Name())-len(fileExtension)]

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

func LoadFile(file entity.File) (string, error) {
	f, err := os.Open(file.GetFullPath())
	if err != nil {
		return "", err
	}
	defer func() error {
		if err = f.Close(); err != nil {
			return err
		}
		return nil
	}()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func CreateFolder(folderPath string) error {
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(filePath string, data string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() error {
		if err = f.Close(); err != nil {
			return err
		}
		return nil
	}()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
