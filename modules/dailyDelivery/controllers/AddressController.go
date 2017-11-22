package controllers

import (
	"encoding/json"
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
type addressController struct {
	mSession *database.MongoSession
}

/**
 *
 */
func AddressController(mongoSession *database.MongoSession) *addressController {
	return &addressController{mongoSession}
}

/**
 *
 */
func (uc *addressController) GetAddress(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	addressDao := dao.NewAddressDao()

	var waitGroup sync.WaitGroup
	address, _ := addressDao.GetAddress(&waitGroup, uc.mSession)
	waitGroup.Wait()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	addrJson, _ := json.Marshal(address)
	fmt.Fprintf(w, "%s", addrJson)
}
