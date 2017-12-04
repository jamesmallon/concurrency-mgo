package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"userv/modules/dailyDelivery"
)

func main() {
	router := httprouter.New()

	dailyDelivery.RouteRegister(router)

	server := &http.Server{Addr: ":3000", Handler: router}
	// server.SetKeepAlivesEnabled(false) // setting keepalive to false
	log.Fatal(server.ListenAndServe())
}
