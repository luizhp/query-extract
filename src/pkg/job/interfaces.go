package job

import (
	"database/sql"

	"github.com/luizhp/query-extract/internal/entity"
	"github.com/luizhp/query-extract/internal/infra/database"
)

type JobInterface interface {
	GetDB() *sql.DB
	GetDBInstance() database.DBInstance
	GetFile() entity.File
	GetResult() entity.Result
	GetOutputFolder() string
	Extract() error
	Dump() error
}
