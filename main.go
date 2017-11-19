package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"userv/commons/database"
	"userv/modules/delivery"
)

func main() {
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
	mongoSession := database.ConnMongo()
	delivery.RouteRegister(router, mongoSession)

	server := &http.Server{Addr: ":3000", Handler: router}
	// server.SetKeepAlivesEnabled(false) // setting keepalive to false
	log.Fatal(server.ListenAndServe())
}
