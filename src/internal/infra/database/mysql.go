package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DBVENDOR = "mysql"

// MySQLInstance implements the DBInstance interface
type MySQLInstance struct {
	db *sql.DB
}

// NewMySQLInstance creates a new MySQLInstance and establishes a connection
func NewMySQLInstance(dsn string) (*MySQLInstance, error) {
	db, err := sql.Open(DBVENDOR, dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("🔗 %s connection established\n", DBVENDOR)
	return &MySQLInstance{db: db}, nil
}

// GetDB returns the underlying *sql.DB instance
func (m *MySQLInstance) GetDB() *sql.DB {
	return m.db
}

// Close closes the database connection
func (m *MySQLInstance) Close() error {
	log.Printf("🔒 Closing %s connection\n", DBVENDOR)
	return m.db.Close()
}

// GetDBVendor returns the database vendor
func (m *MySQLInstance) GetDBVendor() string {
	return DBVENDOR
}
