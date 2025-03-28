package entity

import "time"

type FileInterface interface {
	GetName() string
	GetPath() string
	GetExtension() string
	GetFilename() string
	GetFullPath() string
}

type ResultInterface interface {
	GetColumns() []Column
	GetRows() []map[string]interface{}
	GetTotalRows() int
	GetStartedAt() time.Time
	GetFinishedAt() time.Time
	GetDuration() time.Duration
}

type ColumnInterface interface {
	GetPosition() int
	GetName() string
	GetDatabaseTypeName() string
	GetScanType() interface{}
	GetLength() int64
	GetPrecision() int64
	GetScale() int64
	GetNullable() bool
}
