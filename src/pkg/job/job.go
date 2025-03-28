package job

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/luizhp/query-extract/internal/entity"
	"github.com/luizhp/query-extract/internal/infra/csv"
	"github.com/luizhp/query-extract/internal/infra/filesystem"
)

type Job struct {
	DB           *sql.DB
	File         entity.File
	Result       entity.Result
	OutputFolder string
}

func NewJob(db *sql.DB, file entity.File, outputFolder string) *Job {
	return &Job{
		DB:           db,
		File:         file,
		Result:       entity.Result{},
		OutputFolder: outputFolder,
	}
}

func (b *Job) GetDB() *sql.DB {
	return b.DB
}

func (b *Job) GetFile() entity.File {
	return b.File
}

func (b *Job) GetResult() entity.Result {
	return b.Result
}

func (b *Job) GetOutputFolder() string {
	return b.OutputFolder
}

func (b *Job) Extract() error {

	var startedAt, finishedAt time.Time
	startedAt = time.Now()

	defer func() {
		finishedAt = time.Now()
		log.Printf("🌟 [%s] Data extraction with sucess. It took %v\n", b.File.GetName(), finishedAt.Sub(startedAt))
	}()

	// Load query from file
	query, err := filesystem.LoadFile(b.File)
	if err != nil {
		return err
	}
	log.Printf("📄 [%s] Query loaded\n", b.File.GetName())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Execute query
	log.Printf("🏃 [%s] Executing query\n", b.File.GetName())
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
		columns = append(columns, *entity.NewColumn(i, columnType.Name(), columnType.DatabaseTypeName(), columnType.ScanType(), length, precision, scale, nullable))
	}

	// Get Rows
	var rowsData []map[string]interface{}
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
		rowData := make(map[string]interface{})
		for i, column := range columns {
			val := columnsPointers[i].(*interface{})
			rowData[column.GetName()] = *val
		}
		rowsData = append(rowsData, rowData)
	}
	b.Result = *entity.NewResult(columns, rowsData, startedAt, time.Now())

	return nil

}

func (b *Job) Dump(format string) error {

	var startedAt, finishedAt time.Time
	startedAt = time.Now()

	defer func() {
		finishedAt = time.Now()
		log.Printf("📦 [%s] Dump generated with sucess. It took %v\n", b.File.GetName(), finishedAt.Sub(startedAt))
	}()

	log.Printf("📦 [%s] Dumping %d rows\n", b.File.GetName(), b.Result.GetTotalRows())

	var buffer bytes.Buffer
	switch format {
	case "csv":
		buffer.WriteString(csv.Header(b.Result.GetColumnsName()))
		buffer.WriteString(csv.Detail(b.Result.GetColumnsName(), b.Result.GetRows()))
	default:
		return fmt.Errorf("☠️ Error: Format %s not supported", format)
	}

	if err := filesystem.CreateFolder(b.OutputFolder); err != nil {
		return err
	}

	if err := filesystem.WriteFile(b.OutputFolder+"/"+b.File.GetName()+".csv", buffer.String()); err != nil {
		return err
	}

	// fmt.Println(buffer.String())

	return nil
}
