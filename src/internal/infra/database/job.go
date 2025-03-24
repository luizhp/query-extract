package database

import (
	"bytes"
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
	Results      entity.Results
	OutputFolder string
}

func NewJob(db *sql.DB, file entity.File, outputFolder string) *Job {
	return &Job{
		DB:           db,
		File:         file,
		Results:      entity.Results{},
		OutputFolder: outputFolder,
	}
}

func (b *Job) GetDB() *sql.DB {
	return b.DB
}

func (b *Job) GetFile() entity.File {
	return b.File
}

func (b *Job) GetResults() entity.Results {
	return b.Results
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

	// Execute query
	rows, err := b.GetDB().Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	log.Printf("🏃 [%s] Executing query\n", b.File.GetName())

	// Get Column Names
	columns, err := rows.Columns()
	if err != nil {
		return err
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
			rowData[column] = *val
		}
		rowsData = append(rowsData, rowData)
	}
	b.Results = *entity.NewResults(columns, rowsData, startedAt, time.Now())

	return nil

}

func (b *Job) Dump(format string) error {

	var startedAt, finishedAt time.Time
	startedAt = time.Now()

	defer func() {
		finishedAt = time.Now()
		log.Printf("📦 [%s] Dump generated with sucess. It took %v\n", b.File.GetName(), finishedAt.Sub(startedAt))
	}()

	log.Printf("📦 [%s] Dumping %d rows\n", b.File.GetName(), b.Results.GetTotalRows())

	var buffer bytes.Buffer
	switch format {
	case "csv":
		buffer.WriteString(csv.Header(b.Results.GetColumns()))
		buffer.WriteString(csv.Detail(b.Results.GetColumns(), b.Results.GetRows()))
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
