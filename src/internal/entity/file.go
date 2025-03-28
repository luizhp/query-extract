package entity

import "fmt"

type File struct {
	name      string
	extension string
	path      string
}

func NewFile(name, extension, path string) *File {
	return &File{
		name:      name,
		extension: extension,
		path:      path,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetPath() string {
	return f.path
}

func (f *File) GetExtension() string {
	return f.extension
}

func (f *File) GetFilename() string {
	return f.name + f.extension
}

func (f *File) GetFullPath() string {
	return fmt.Sprintf("%s/%s.%s", f.path, f.name, f.extension)
}
