package delivery

import (
	"github.com/julienschmidt/httprouter"
	//"net/http"
	"userv/modules/delivery/controllers"
)

func RouteRegister(route *httprouter.Router) {
	carController := controllers.CarController()

	route.GET("/", carController.GetCar)
}
