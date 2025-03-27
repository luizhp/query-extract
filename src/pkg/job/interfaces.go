package job

import (
	"database/sql"

	"github.com/luizhp/query-extract/internal/entity"
)

type JobInterface interface {
	GetDB() *sql.DB
	GetFile() entity.File
	GetResult() entity.Result
	GetOutputFolder() string
	Extract() error
	Dump() error
}
