package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
	"userv/commons/database"
	"userv/modules/dailyDelivery/dao"
)

/**
 *
 */
type deliveryController struct {
	mSession *database.MongoSession
}

/**
 *
 */
func DeliveryController(mongoSession *database.MongoSession) *deliveryController {
	return &deliveryController{mongoSession}
}

/**
 *
 */
func (uc *deliveryController) GetDelivery(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deliveryDao := dao.NewDeliveryDao()
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go deliveryDao.RunQuery(&waitGroup, uc.mSession)

	waitGroup.Wait()
	fmt.Println("Query Completed")
}
