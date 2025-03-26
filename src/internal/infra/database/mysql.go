package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLInstance implements the DBInstance interface
type MySQLInstance struct {
	db *sql.DB
}

// NewMySQLInstance creates a new MySQLInstance and establishes a connection
func NewMySQLInstance(dsn string) (*MySQLInstance, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("🔗 MySQL connection established")
	return &MySQLInstance{db: db}, nil
}

// GetDB returns the underlying *sql.DB instance
func (m *MySQLInstance) GetDB() *sql.DB {
	return m.db
}

// Close closes the database connection
func (m *MySQLInstance) Close() error {
	log.Println("🔒 Closing MySQL connection")
	return m.db.Close()
}
