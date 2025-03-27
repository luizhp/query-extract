package database

import (
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

// MSSQLInstance implements the DBInstance interface
type MSSQLInstance struct {
	vendor string
	db     *sql.DB
}

// NewMSSQLInstance creates a new MSSQLInstance and establishes a connection
func NewMSSQLInstance(dsn string) (*MSSQLInstance, error) {
	var vendor = "sqlserver"
	db, err := sql.Open(vendor, dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("🔗 %s connection established\n", vendor)
	return &MSSQLInstance{vendor: vendor, db: db}, nil
}

// GetDB returns the underlying *sql.DB instance
func (m *MSSQLInstance) GetDB() *sql.DB {
	return m.db
}

// GetDBVendor returns the database vendor
func (m *MSSQLInstance) GetDBVendor() string {
	return m.vendor
}

// Close closes the database connection
func (m *MSSQLInstance) Close() error {
	log.Printf("🔒 Closing %s connection\n", m.vendor)
	return m.db.Close()
}
