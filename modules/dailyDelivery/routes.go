package dailyDelivery

import (
	"github.com/julienschmidt/httprouter"
	//"userv/commons/cache"
	"userv/commons/database"
	"userv/modules/dailyDelivery/controllers"
)

func RouteRegister(route *httprouter.Router, mongoSession *database.MongoSession) {
	deliveryController := controllers.DeliveryController(mongoSession)

	/*
		clear; curl -X GET 'http://127.0.0.1:3000/delivery' \
		--data-binary '{"date": "2017-10-15"}'
	*/
	// ab -qrk -c 500 -n 1000 "http://127.0.0.1:3000/delivery"
	route.GET("/delivery", deliveryController.GetDelivery)
}
