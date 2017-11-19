package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
	"userv/modules/delivery/dao"
)

/**
 *
 */
type carController struct {
}

/**
 *
 */
func CarController() *carController {
	return new(carController)
}

/**
 *
 */
func (uc *carController) GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carDao := dao.NewCarDao()
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go carDao.RunQuery(&waitGroup)

	waitGroup.Wait()
	fmt.Println("Query Completed")
}
