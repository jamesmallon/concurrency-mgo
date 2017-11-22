package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
	"userv/commons/cache"
	"userv/commons/database"
	//"userv/modules/dailyDelivery/dao"
	//"userv/modules/dailyDelivery/models"
)

/**
 *
 */
type deliveryController struct {
}

/**
 *
 */
func DeliveryController() *deliveryController {
	return &deliveryController{}
}

/**
 *
 */
func (uc *deliveryController) GetDelivery(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//deliveryDao := dao.NewDeliveryDao()

	var waitGroup sync.WaitGroup
	address, _ := database.GetThat(&waitGroup)
	waitGroup.Wait()

	var waitGroupA sync.WaitGroup
	cache.Set("thiIsAKey", "keysValue", &waitGroupA)
	waitGroupA.Wait()

	var waitGroupB sync.WaitGroup
	fmt.Println(cache.Get("thiIsAKey", &waitGroupB))
	waitGroupB.Wait()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	dlryJson, _ := json.Marshal(address)
	fmt.Fprintf(w, "%s", dlryJson)
}
