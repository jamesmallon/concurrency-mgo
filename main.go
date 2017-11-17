package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sync"
	"userv/commons/database"
	"userv/modules/delivery"
)

func main() {
	// Get Mongodb instance connection
	mongoSession := database.ConnMongo()

	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup

	// Perform 1 concurrent queries against the database.
	waitGroup.Add(1)
	go mongoSession.RunQuery(&waitGroup)

	// Wait for all the queries to complete.
	waitGroup.Wait()
	fmt.Println("All Queries Completed")

	router := httprouter.New()

	// transport := &http.Transport{
	// 	Proxy: http.ProxyFromEnvironment,
	// 	Dial: (&net.Dialer{
	// 		Timeout:   30 * time.Second,
	// 		KeepAlive: time.Minute,
	// 	}).Dial,
	// 	TLSHandshakeTimeout: 10 * time.Second,
	// }
	//
	// router := &http.Client{
	// 	Transport: transport,
	// }

	delivery.RouteRegister(router)

	server := &http.Server{Addr: ":3000", Handler: router}
	// server.SetKeepAlivesEnabled(false) // setting keepalive to false
	log.Fatal(server.ListenAndServe())
}
