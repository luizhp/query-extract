package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/luizhp/query-extract/internal/infra/database"
	"github.com/luizhp/query-extract/internal/infra/filesystem"

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
	mysqlDB, err := sql.Open("mysql", "appuser:apppassword@tcp(localhost:3306)/appdb?parseTime=false")
	if err != nil {
		log.Printf("☠️ Error: %v\n", err)
		os.Exit(1)
	} else {
		log.Printf("🔗 Open mysql connection")
	}
	mysqlDB.SetConnMaxLifetime(time.Minute * 20)
	defer mysqlDB.Close()

	// Process each job
	for _, queryFile := range queriesCollection {
		job := database.NewJob(mysqlDB, queryFile, outputFolder)
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
