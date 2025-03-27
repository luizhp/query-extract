package database

import (
	"database/sql"
)

type DBInstance interface {
	GetDB() *sql.DB
	GetDBVendor() string
	Close() error
}
