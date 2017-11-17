package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
 *
 */
type carController struct{}

/**
 *
 */
func CarController() *carController {
	return new(carController)
}

/**
 *
 */
func (uc carController) GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("hey you!")
}
