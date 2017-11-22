package dailyDelivery

import (
	"github.com/julienschmidt/httprouter"
	//"userv/commons/database"
	"userv/modules/dailyDelivery/controllers"
)

func RouteRegister(route *httprouter.Router) {
	//addressController := controllers.AddressController(mongoSession)
	deliveryController := controllers.DeliveryController()

	/*
		clear; curl -X GET 'http://127.0.0.1:3000/address' \
		--data-binary '{"date": "2017-10-15"}'
	*/
	// ab -qrk -c 500 -n 1000 "http://127.0.0.1:3000/address"
	//route.GET("/address", addressController.GetAddress)

	/*
		clear; curl -X GET 'http://127.0.0.1:3000/delivery' \
		--data-binary '{"date": "2017-10-15"}'
	*/
	// ab -qrk -c 500 -n 1000 "http://127.0.0.1:3000/delivery"
	route.GET("/delivery", deliveryController.GetDelivery)
}
