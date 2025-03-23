package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/luizhp/query-extract/internal/entity"
)

type Job struct {
	DB      *sql.DB
	File    entity.File
	Results []entity.Results
}

func NewJob(db *sql.DB, file entity.File) *Job {
	return &Job{
		DB:      db,
		File:    file,
		Results: []entity.Results{},
	}
}

func (b *Job) GetDB() *sql.DB {
	return b.DB
}

func (b *Job) GetFile() entity.File {
	return b.File
}

func (b *Job) GetResults() []entity.Results {
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

	// Execute query

	// Save results
	// // Get Column Names
	// // Get Rows

	// b.Results = entity.NewResults(columns, rows, startedAt, finishedAt)

	return nil

}
