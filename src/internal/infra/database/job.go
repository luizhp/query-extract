package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/luizhp/query-extract/internal/entity"
	"github.com/luizhp/query-extract/internal/infra/filesystem"
)

type Job struct {
	DB      *sql.DB
	File    entity.File
	Results entity.Results
}

func NewJob(db *sql.DB, file entity.File) *Job {
	return &Job{
		DB:      db,
		File:    file,
		Results: entity.Results{},
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

func (b *Job) Process() error {

	var startedAt, finishedAt time.Time
	startedAt = time.Now()

	defer func() {
		finishedAt = time.Now()
		log.Printf("🌟 Job finished with sucess. It took %v\n", finishedAt.Sub(startedAt))
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

	// Save results
	// // Get Column Names
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// // Get Rows
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
