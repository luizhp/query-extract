package main

import (
	"log"
	"os"
	"time"

	"github.com/luizhp/query-extract/internal/infra/database"
	"github.com/luizhp/query-extract/internal/infra/filesystem"
	"github.com/luizhp/query-extract/pkg/job"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	startTimerAll := time.Now()
	log.Printf("🚀 Start\n")
	defer func() {
		log.Printf("🏆 Finished with sucess. It took %v\n", time.Since(startTimerAll))
		os.Exit(0)
	}()

	queriesFolder := "/home/user/Documents/dev/luizhp/QueryExtract/data/queries"
	queriesExtension := "sql"
	outputFolder := "/home/user/Documents/dev/luizhp/QueryExtract/data/output"

	// Get list of query files available
	queriesCollection, err := filesystem.ListFolder(queriesFolder, queriesExtension)
	if err != nil {
		log.Printf("☠️ Error: %v\n", err)
		os.Exit(1)
	}
	if len(queriesCollection) == 0 {
		log.Printf("☠️ Error: No queries found at %v\n", queriesFolder)
		os.Exit(1)
	} else {
		log.Printf("📁 Found %v queries\n", len(queriesCollection))
	}

	// Get Target DB connection
	// var db database.DBInstance
	// dsn := "sqlserver://sa:StrongPassword123@localhost?database=queryextract"
	// db, err = database.NewMSSQLInstance(dsn)
	// if err != nil {
	// 	log.Printf("☠️ Error: %v\n", err)
	// 	os.Exit(1)
	// }
	// db.GetDB().SetConnMaxLifetime(time.Minute * 20)
	// defer db.Close()

	// MySQL 8
	var db database.DBInstance
	dsn := "appuser:apppassword@tcp(localhost:3306)/appdb?parseTime=false"
	db, err = database.NewMySQLInstance(dsn)
	if err != nil {
		log.Printf("☠️ Error: %v\n", err)
		os.Exit(1)
	}
	db.GetDB().SetConnMaxLifetime(time.Minute * 20)
	defer db.Close()

	// Process each job
	for _, queryFile := range queriesCollection {
		job := job.NewJob(db, queryFile, outputFolder)
		if err := job.Extract(); err != nil {
			log.Printf("☠️ Error: %v\n", err)
			os.Exit(1)
		}
		// Dump Result to a csv file
		if err := job.Dump("csv"); err != nil {
			log.Printf("☠️ Error: %v\n", err)
			os.Exit(1)
		}
	}

}
