package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sync"
	"userv/commons/database"
)

func main() {
	router := httprouter.New()

	mongoSession := database.ConnMongo()

	// clear; curl -X GET 'http://127.0.0.1:3000/'
	// clear; ab -qrk -c 10 -n 50 "http://127.0.0.1:3000/"
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// Create a wait group to manage the goroutines.
		var waitGroup sync.WaitGroup

		// Perform 10 concurrent queries against the database.
		waitGroup.Add(1)
		go database.RunQuery(&waitGroup, mongoSession)

		// Wait for all the queries to complete.
		waitGroup.Wait()
	})

	server := &http.Server{Addr: ":3000", Handler: router}
	server.SetKeepAlivesEnabled(false) // setting keepalive to false
	log.Fatal(server.ListenAndServe())
}
