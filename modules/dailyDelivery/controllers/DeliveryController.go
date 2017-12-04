package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	//"sync"
	"userv/commons/cache"
	"userv/commons/database"
	"userv/modules/dailyDelivery/dao"
)

/**
 *
 */
type deliveryController struct {
	mSession *database.MongoSession
	rClient  *cache.RedisClient
}

/**
 *
 */
func DeliveryController(mongoSession *database.MongoSession, redisClient *cache.RedisClient) *deliveryController {
	return &deliveryController{mongoSession, redisClient}
}

/**
 * @method GetDelivery
 */
func (uc *deliveryController) GetDelivery(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deliveryDao := dao.NewDeliveryDao()
	delivery, _ := deliveryDao.GetDelivery(uc.mSession)
	delivery, _ = deliveryDao.IncrementField(uc.mSession, "sussDlry", delivery)
	fmt.Println(delivery)

	deliveryCacheDao := dao.NewDeliveryCacheDao()
	deliveryCacheDao.SettingKey("trust", "+1000000 4U", uc.rClient)
	fmt.Println(deliveryCacheDao.GettingKey("trust", uc.rClient))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	dlryJson, _ := json.Marshal(delivery)
	fmt.Fprintf(w, "%s", dlryJson)
}
