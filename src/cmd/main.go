package main

import (
	"log"
	"os"
	"time"

	"github.com/luizhp/query-extract/internal/infra/filesystem"
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
	// outputFolder := "/home/user/Documents/dev/luizhp/QueryExtract/data/output"

	// Get list of query files available
	queries, err := filesystem.ListFolder(queriesFolder, queriesExtension)
	if err != nil {
		log.Printf("☠️ Error: %v\n", err)
		os.Exit(1)
	}
	if len(queries) == 0 {
		log.Printf("☠️ Error: No queries found at %v\n", queriesFolder)
		os.Exit(1)
	}

	// Get Target DB connection

	// Loop through each query file
	// 	- Read the query file
	// 	- Execute the query
	// 	- Dump Result to a csv file

}
