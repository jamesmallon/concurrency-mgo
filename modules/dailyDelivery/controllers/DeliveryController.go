package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
	"userv/commons/database"
	"userv/modules/dailyDelivery/dao"
	//"userv/modules/dailyDelivery/models"
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
	delivery, _ := deliveryDao.GetDelivery(&waitGroup, uc.mSession)
	waitGroup.Wait()
	delivery, _ = deliveryDao.IncrementField(&waitGroup, uc.mSession, "sussDlry", delivery)
	waitGroup.Wait()
	fmt.Println(delivery)
}
