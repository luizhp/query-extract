package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLInstance implements the DBInstance interface
type MySQLInstance struct {
	vendor string
	db     *sql.DB
}

// NewMySQLInstance creates a new MySQLInstance and establishes a connection
func NewMySQLInstance(dsn string) (*MySQLInstance, error) {
	var vendor = "mysql"
	db, err := sql.Open(vendor, dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("🔗 %s connection established\n", vendor)
	return &MySQLInstance{vendor: vendor, db: db}, nil
}

// GetDB returns the underlying *sql.DB instance
func (m *MySQLInstance) GetDB() *sql.DB {
	return m.db
}

// GetDBVendor returns the database vendor
func (m *MySQLInstance) GetDBVendor() string {
	return m.vendor
}

// Close closes the database connection
func (m *MySQLInstance) Close() error {
	log.Printf("🔒 Closing %s connection\n", m.vendor)
	return m.db.Close()
}
