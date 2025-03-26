package database

import (
	"database/sql"
)

type DBInstance interface {
	GetDBVendor() string
	GetDB() *sql.DB
	Close() error
}
