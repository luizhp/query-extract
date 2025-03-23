package database

import (
	"database/sql"

	"github.com/luizhp/query-extract/internal/entity"
)

type JobInterface interface {
	GetDB() *sql.DB
	GetFile() entity.File
	Process() error
}
