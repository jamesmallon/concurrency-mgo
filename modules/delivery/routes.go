package delivery

import (
	"github.com/julienschmidt/httprouter"
	//"net/http"
	"userv/modules/delivery/controllers"
)

func RouteRegister(route *httprouter.Router) {
	carController := controllers.CarController()

	// ab -qrk -c 1 -n 5 "http://127.0.0.1:3000/"
	route.GET("/", carController.GetCar)
}
