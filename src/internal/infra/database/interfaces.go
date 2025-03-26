package database

import (
	"database/sql"
)

type DBInstance interface {
	GetDB() *sql.DB
	Close() error
}
