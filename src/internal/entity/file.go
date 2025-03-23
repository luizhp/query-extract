package entity

import "fmt"

type File struct {
	Name      string
	Extension string
	Path      string
}

func NewFile(name, extension, path string) *File {
	return &File{
		Name:      name,
		Extension: extension,
		Path:      path,
	}
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetPath() string {
	return f.Path
}

func (f *File) GetExtension() string {
	return f.Extension
}

func (f *File) GetFilename() string {
	return f.Name + f.Extension
}

func (f *File) GetFullPath() string {
	return fmt.Sprintf("%s/%s%s", f.Path, f.Name, f.Extension)
}
