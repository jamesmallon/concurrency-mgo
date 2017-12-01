package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

	delivery, _ := deliveryDao.GetDelivery(uc.mSession)
	delivery, _ = deliveryDao.IncrementField(uc.mSession, "sussDlry", delivery)
	fmt.Println(delivery)

	deliveryCacheDao := dao.NewDeliveryCacheDao()
	deliveryCacheDao.SettingKey("mingal")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	dlryJson, _ := json.Marshal(delivery)
	fmt.Fprintf(w, "%s", dlryJson)
}
