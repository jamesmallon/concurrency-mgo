package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"userv/commons/database"
	"userv/modules/dailyDelivery"
)

func main() {
	router := httprouter.New()

	mongoSession := database.ConnMongo()
	dailyDelivery.RouteRegister(router, mongoSession)

	server := &http.Server{Addr: ":3000", Handler: router}
	// server.SetKeepAlivesEnabled(false) // setting keepalive to false
	log.Fatal(server.ListenAndServe())
}
