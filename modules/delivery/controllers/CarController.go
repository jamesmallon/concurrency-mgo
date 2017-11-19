package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
	"userv/commons/database"
	"userv/modules/delivery/dao"
)

/**
 *
 */
type carController struct {
	mSession *database.MongoSession
}

/**
 *
 */
func CarController(mongoSession *database.MongoSession) *carController {
	return &carController{mongoSession}
}

/**
 *
 */
func (uc *carController) GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carDao := dao.NewCarDao()
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go carDao.RunQuery(&waitGroup, uc.mSession)

	waitGroup.Wait()
	fmt.Println("Query Completed")
}
