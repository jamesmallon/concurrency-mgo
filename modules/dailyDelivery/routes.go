package dailyDelivery

import (
	"github.com/julienschmidt/httprouter"
	//"net/http"
	"userv/commons/database"
	"userv/modules/dailyDelivery/controllers"
)

func RouteRegister(route *httprouter.Router, mongoSession *database.MongoSession) {
	addressController := controllers.AddressController(mongoSession)
	deliveryController := controllers.DeliveryController(mongoSession)

	/*
		clear; curl -X GET 'http://127.0.0.1:3000/address' \
		--data-binary '{"domain": "carrierexpress.com.br/"}'
	*/
	// ab -qrk -c 1 -n 5 "http://127.0.0.1:3000/address"
	route.GET("/address", addressController.GetAddress)

	/*
		clear; curl -X GET 'http://127.0.0.1:3000/delivery' \
		--data-binary '{"domain": "carrierexpress.com.br/"}'
	*/
	// ab -qrk -c 1 -n 5 "http://127.0.0.1:3000/delivery"
	route.GET("/delivery", deliveryController.GetDelivery)
}
