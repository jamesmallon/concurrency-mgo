package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"userv/modules/dailyDelivery"
)

func main() {
	//redisAddr := flag.String("redis-addr", "and", "then")
	//redisKey := flag.String("redis-key", "so", "what")
	//flag.Parse()

	//fmt.Println(*redisAddr)
	//fmt.Println(*redisKey)

	router := httprouter.New()

	dailyDelivery.RouteRegister(router)

	server := &http.Server{Addr: ":3000", Handler: router}
	// server.SetKeepAlivesEnabled(false) // setting keepalive to false
	log.Fatal(server.ListenAndServe())
}
