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
	b.result = *entity.NewResult(columns, rowsData, startedAt, time.Now())

	return nil

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
		buffer.WriteString(csv.Header(b.result.GetColumnsName()))
		buffer.WriteString(csv.Detail(b.result.GetColumnsName(), b.result.GetRows()))
	default:
		return fmt.Errorf("☠️ Error: Format %s not supported", format)
	}

	if err := filesystem.CreateFolder(b.outputFolder); err != nil {
		return err
	}

	if err := filesystem.WriteFile(b.outputFolder+"/"+b.file.GetName()+".csv", buffer.String()); err != nil {
		return err
	}

	// fmt.Println(buffer.String())

	return nil
}
