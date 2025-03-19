package main

import (
	"log"
	"os"
	"time"
)

func main() {

	startTimerAll := time.Now()
	log.Printf("🚀 Start\n")
	defer func() {
		log.Printf("🏆 Finished with sucess. It took %v\n", time.Since(startTimerAll))
		os.Exit(0)
	}()

	// Get list of query files available

	// Get Target DB connection

	// Loop through each query file
	// 	- Read the query file
	// 	- Execute the query
	// 	- Dump Result to a csv file

}
