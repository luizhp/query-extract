package entity

type FileInterface interface {
	GetName() string
	GetPath() string
	GetExtension() string
	GetFilename() string
}

type ResultsInterface interface {
	GetColumns() []string
	GetRows() []map[string]interface{}
	GetTotalRows() int
	GetStartedAt() string
	GetFinishedAt() string
	GetDuration() string
}
