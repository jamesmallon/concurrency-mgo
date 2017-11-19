package delivery

import (
	"github.com/julienschmidt/httprouter"
	//"net/http"
	"userv/commons/database"
	"userv/modules/delivery/controllers"
)

func RouteRegister(route *httprouter.Router, mongoSession *database.MongoSession) {
	carController := controllers.CarController(mongoSession)

	// ab -qrk -c 1 -n 5 "http://127.0.0.1:3000/"
	route.GET("/", carController.GetCar)
}
