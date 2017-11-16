package main

import (
	"fmt"
	"sync"
	"userv/commons/database"
)

func main() {
	// Get Mongodb instance connection
	mongoSession := database.ConnMongo()

	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup

	// Perform 10 concurrent queries against the database.
	waitGroup.Add(1)
	go mongoSession.RunQuery(&waitGroup)

	// Wait for all the queries to complete.
	waitGroup.Wait()
	fmt.Println("All Queries Completed")

}
