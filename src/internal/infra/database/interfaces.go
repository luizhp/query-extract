package database

import (
	"database/sql"

	"github.com/luizhp/query-extract/internal/entity"
)

type JobInterface interface {
	GetDB() *sql.DB
	GetFile() entity.File
	GetResults() entity.Results
	GetOutputFolder() string
	Extract() error
	Dump() error
}
