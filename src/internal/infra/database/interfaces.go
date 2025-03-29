package database

import (
	"database/sql"

	"github.com/luizhp/query-extract/internal/entity"
)

type DBInstance interface {
	GetDB() *sql.DB
	GetDBVendor() string
	Close() error
	Convert(dataType entity.Column, dataValue *interface{}) (string, error)
}
