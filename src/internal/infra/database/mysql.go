package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luizhp/query-extract/internal/entity"
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

func (m *MySQLInstance) Convert(dataType entity.Column, dataValue *interface{}) (string, error) {
	var convertedValue string = ""

	if dataValue == nil || dataType.GetScanType() == nil {
		return convertedValue, nil
	}

	switch dataType.GetScanType().Kind() {
	// Integer
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		convertedValue = fmt.Sprintf("%d", *dataValue)
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		convertedValue = fmt.Sprintf("%d", *dataValue)
	// // Decimal
	// case reflect.Slice:
	// 	switch dataType.GetScanType().Elem().Kind() {
	// 	case reflect.Uint8:
	// 		switch dataType.GetDatabaseTypeName() {
	// 		case "DECIMAL":
	// 			convertedValue = string((*dataValue).([]uint8))
	// 		case "IMAGE":
	// 			convertedValue = ""
	// 		case "UNIQUEIDENTIFIER":
	// 			convertedValue = ""
	// 		case "GEOGRAPHY", "GEOMETRY", "HIERARCHYID":
	// 			convertedValue = ""
	// 		default:
	// 			convertedValue = string((*dataValue).([]uint8))
	// 		}
	// 	default:
	// 		convertedValue = string((*dataValue).([]uint8))
	// 	}
	// // Float
	// case reflect.Float64, reflect.Float32:
	// 	switch dataType.GetDatabaseTypeName() {
	// 	case "FLOAT":
	// 		convertedValue = fmt.Sprintf("%g", *dataValue)
	// 	case "REAL":
	// 		convertedValue = fmt.Sprintf("%f", *dataValue)
	// 	default:
	// 		convertedValue = fmt.Sprintf("%f", *dataValue)
	// 	}
	// 	convertedValue = strings.TrimRight(convertedValue, "0")
	// 	// String
	// case reflect.String:
	// 	switch dataType.GetDatabaseTypeName() {
	// 	case "CHAR":
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	case "VARCHAR":
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	case "TEXT":
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	case "NCHAR":
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	case "NVARCHAR":
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	case "NTEXT":
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	default:
	// 		convertedValue = fmt.Sprintf("%s", *dataValue)
	// 	}
	// // Bool
	// case reflect.Bool:
	// 	convertedValue = fmt.Sprintf("%t", *dataValue)
	// Struct
	case reflect.Struct:
		switch dataType.GetScanType() {
		case reflect.TypeOf(sql.NullInt64{}), reflect.TypeOf(sql.NullInt32{}), reflect.TypeOf(sql.NullInt16{}):
			convertedValue = fmt.Sprintf("%d", *dataValue)
		// case reflect.TypeOf(time.Time{}):
		// 	switch dataType.GetDatabaseTypeName() {
		// 	case "DATE":
		// 		convertedValue = (*dataValue).(time.Time).Format("2006-01-02 ")
		// 	case "TIME":
		// 		convertedValue = (*dataValue).(time.Time).Format("15:04:05")
		// 	case "DATETIME", "SMALLDATETIME", "DATETIME2":
		// 		convertedValue = (*dataValue).(time.Time).Format("2006-01-02 15:04:05.000 ")
		// 	case "DATETIMEOFFSET":
		// 		convertedValue = (*dataValue).(time.Time).Format("2006-01-02 15:04:05.000 -0700")
		// 	default:
		// 		convertedValue = (*dataValue).(time.Time).Format("2006-01-02 15:04:05.000 ")
		// 	}
		case reflect.TypeOf(sql.NullString{}):
			switch dataType.GetDatabaseTypeName() {
			case "DECIMAL":
				convertedValue = fmt.Sprintf("%s", *dataValue)
			case "CHAR":
				convertedValue = fmt.Sprintf("%s", *dataValue)
			case "VARCHAR":
				convertedValue = fmt.Sprintf("%s", *dataValue)
			case "TEXT":
				convertedValue = fmt.Sprintf("%s", *dataValue)
			default:
				fmt.Printf("coluna: %s - formato: %s - kind: %s - dbformat: %s\n", dataType.GetName(), dataType.GetScanType(), dataType.GetScanType().Kind(), dataType.GetDatabaseTypeName())
				convertedValue = fmt.Sprintf("%v", *dataValue)
			}
		case reflect.TypeOf(sql.NullFloat64{}):
			switch dataType.GetDatabaseTypeName() {
			case "FLOAT":
				convertedValue = fmt.Sprintf("%g", *dataValue)
			case "DOUBLE":
				convertedValue = fmt.Sprintf("%g", *dataValue)
			default:
				fmt.Printf("coluna: %s - formato: %s - kind: %s - dbformat: %s\n", dataType.GetName(), dataType.GetScanType(), dataType.GetScanType().Kind(), dataType.GetDatabaseTypeName())
				convertedValue = fmt.Sprintf("%v", *dataValue)
			}
		default:
			fmt.Printf("coluna: %s - formato: %s - kind: %s - dbformat: %s\n", dataType.GetName(), dataType.GetScanType(), dataType.GetScanType().Kind(), dataType.GetDatabaseTypeName())
			convertedValue = fmt.Sprintf("%v", *dataValue)
		}
	default:
		switch dataType.GetDatabaseTypeName() {
		// case "SQL_VARIANT":
		// 	convertedValue = ""
		default:
			fmt.Printf("coluna: %s - formato: %s - kind: %s - dbformat: %s\n", dataType.GetName(), dataType.GetScanType(), dataType.GetScanType().Kind(), dataType.GetDatabaseTypeName())
			convertedValue = fmt.Sprintf("%v", *dataValue)
		}
	}
	return convertedValue, nil
}
