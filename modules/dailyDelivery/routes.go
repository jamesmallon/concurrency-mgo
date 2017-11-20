package dailyDelivery

import (
	"github.com/julienschmidt/httprouter"
	//"net/http"
	"userv/commons/database"
	"userv/modules/dailyDelivery/controllers"
)

func RouteRegister(route *httprouter.Router, mongoSession *database.MongoSession) {
	//addressController := controllers.AddressController(mongoSession)
	addressController := controllers.AddressController(mongoSession)

	/*
		clear; curl -X GET 'http://127.0.0.1:3000/' \
		--data-binary '{"domain": "carrierexpress.com.br/"}'
	*/
	// ab -qrk -c 1 -n 5 "http://127.0.0.1:3000/"
	route.GET("/", addressController.GetAddress)
}
