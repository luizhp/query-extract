package job

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/luizhp/query-extract/internal/entity"
	"github.com/luizhp/query-extract/internal/infra/csv"
	"github.com/luizhp/query-extract/internal/infra/filesystem"
	"github.com/luizhp/query-extract/pkg/strutil"
)

type Job struct {
	db           *sql.DB
	file         entity.File
	result       entity.Result
	outputFolder string
}

func NewJob(db *sql.DB, file entity.File, outputFolder string) *Job {
	return &Job{
		db:           db,
		file:         file,
		result:       entity.Result{},
		outputFolder: outputFolder,
	}
}

func (b *Job) GetDB() *sql.DB {
	return b.db
}

func (b *Job) GetFile() entity.File {
	return b.file
}

func (b *Job) GetResult() entity.Result {
	return b.result
}

func (b *Job) GetOutputFolder() string {
	return b.outputFolder
}

func (b *Job) Extract() error {

	var startedAt, finishedAt time.Time
	startedAt = time.Now()

	defer func() {
		finishedAt = time.Now()
		log.Printf("🌟 [%s] Data extraction with sucess. It took %v\n", b.file.GetName(), finishedAt.Sub(startedAt))
	}()

	// Load query from file
	query, err := filesystem.LoadFile(b.file)
	if err != nil {
		return err
	}
	log.Printf("📄 [%s] Query loaded\n", b.file.GetName())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Execute query
	log.Printf("🏃 [%s] Executing query\n", b.file.GetName())
	rows, err := b.GetDB().QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Get Columns' Metadata
	var columns []entity.Column

	columnsTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	for i, columnType := range columnsTypes {
		length, _ := columnType.Length()
		precision, scale, _ := columnType.DecimalSize()
		nullable, _ := columnType.Nullable()
		var rt reflect.Type = columnType.ScanType()
		columns = append(columns, *entity.NewColumn(i, columnType.Name(), columnType.DatabaseTypeName(), rt, length, precision, scale, nullable))
	}

	// Get Rows
	var rowsData []map[string]string
	for rows.Next() {
		columnsData := make([]interface{}, len(columns))
		columnsPointers := make([]interface{}, len(columns))
		for i := range columnsData {
			columnsPointers[i] = &columnsData[i]
		}
		err := rows.Scan(columnsPointers...)
		if err != nil {
			return err
		}
		rowData := make(map[string]string)
		for i, column := range columns {
			val := columnsPointers[i].(*interface{})
			convertedValue, err := b.convert(column, val)
			if err != nil {
				return err
			}
			rowData[column.GetName()] = strutil.RemoveSpecialCodes(convertedValue)
		}
		rowsData = append(rowsData, rowData)
	}
	b.result = *entity.NewResult(columns, rowsData, startedAt, time.Now())

	return nil

}

func (b *Job) convert(dataType entity.Column, dataValue *interface{}) (string, error) {
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
	// Decimal
	case reflect.Slice:
		switch dataType.GetScanType().Elem().Kind() {
		case reflect.Uint8:
			switch dataType.GetDatabaseTypeName() {
			case "DECIMAL":
				convertedValue = string((*dataValue).([]uint8))
			case "IMAGE":
				convertedValue = ""
			case "UNIQUEIDENTIFIER":
				convertedValue = ""
			case "GEOGRAPHY", "GEOMETRY", "HIERARCHYID":
				convertedValue = ""
			default:
				convertedValue = string((*dataValue).([]uint8))
			}
		default:
			convertedValue = string((*dataValue).([]uint8))
		}
	// Float
	case reflect.Float64, reflect.Float32:
		switch dataType.GetDatabaseTypeName() {
		case "FLOAT":
			convertedValue = fmt.Sprintf("%g", *dataValue)
		case "REAL":
			convertedValue = fmt.Sprintf("%f", *dataValue)
		default:
			convertedValue = fmt.Sprintf("%f", *dataValue)
		}
		convertedValue = strings.TrimRight(convertedValue, "0")
		// String
	case reflect.String:
		switch dataType.GetDatabaseTypeName() {
		case "CHAR":
			convertedValue = fmt.Sprintf("%s", *dataValue)
		case "VARCHAR":
			convertedValue = fmt.Sprintf("%s", *dataValue)
		case "TEXT":
			convertedValue = fmt.Sprintf("%s", *dataValue)
		case "NCHAR":
			convertedValue = fmt.Sprintf("%s", *dataValue)
		case "NVARCHAR":
			convertedValue = fmt.Sprintf("%s", *dataValue)
		case "NTEXT":
			convertedValue = fmt.Sprintf("%s", *dataValue)
		default:
			convertedValue = fmt.Sprintf("%s", *dataValue)
		}
	// Bool
	case reflect.Bool:
		convertedValue = fmt.Sprintf("%t", *dataValue)
	// Struct
	case reflect.Struct:
		switch dataType.GetScanType() {
		case reflect.TypeOf(time.Time{}):
			switch dataType.GetDatabaseTypeName() {
			case "DATE":
				convertedValue = (*dataValue).(time.Time).Format("2006-01-02 ")
			case "TIME":
				convertedValue = (*dataValue).(time.Time).Format("15:04:05")
			case "DATETIME", "SMALLDATETIME", "DATETIME2":
				convertedValue = (*dataValue).(time.Time).Format("2006-01-02 15:04:05.000 ")
			case "DATETIMEOFFSET":
				convertedValue = (*dataValue).(time.Time).Format("2006-01-02 15:04:05.000 -0700")
			default:
				convertedValue = (*dataValue).(time.Time).Format("2006-01-02 15:04:05.000 ")
			}
		default:
			convertedValue = fmt.Sprintf("%v", *dataValue)
		}
	default:
		switch dataType.GetDatabaseTypeName() {
		case "SQL_VARIANT":
			convertedValue = ""
		default:
			fmt.Printf("coluna: %s - formato: %s - kind: %s - dbformat: %s\n", dataType.GetName(), dataType.GetScanType(), dataType.GetScanType().Kind(), dataType.GetDatabaseTypeName())
			convertedValue = fmt.Sprintf("%v", *dataValue)
		}
	}
	return convertedValue, nil
}

func (b *Job) Dump(format string) error {

	var startedAt, finishedAt time.Time
	startedAt = time.Now()

	defer func() {
		finishedAt = time.Now()
		log.Printf("📦 [%s] Dump generated with sucess. It took %v\n", b.file.GetName(), finishedAt.Sub(startedAt))
	}()

	log.Printf("📦 [%s] Dumping %d rows\n", b.file.GetName(), b.result.GetTotalRows())

	var buffer bytes.Buffer
	switch format {
	case "csv":
		buffer.WriteString(csv.Generate(b.result.GetColumnsName(), b.result.GetRows()))
	default:
		return fmt.Errorf("☠️ Error: Format %s not supported", format)
	}

	if err := filesystem.CreateFolder(b.outputFolder); err != nil {
		return err
	}

	if err := filesystem.WriteFile(b.outputFolder+"/"+b.file.GetName()+".csv", buffer.String()); err != nil {
		return err
	}

	return nil
}
