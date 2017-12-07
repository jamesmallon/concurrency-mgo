package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	//"sync"
	"github.com/johnthegreenobrien/LogIt"
	"time"
	"userv/commons/cache"
	"userv/commons/database"
	//"userv/commons/log"
	"userv/modules/dailyDelivery/dao"
	"userv/modules/dailyDelivery/models"
)

/**
 *
 */
type deliveryController struct {
	mSession  *database.MongoSession
	rClient   *cache.RedisClient
	dailyColl string
	logging   *logit.SysLog
}

/**
 *
 */
func DeliveryController() *deliveryController {
	return &deliveryController{mSession: database.ConnMongo(), rClient: cache.ConnRedis()}
}

/**
 * Compare length of the daily delivery and addresses collection
 */
func (uc *deliveryController) compareDeliveryAddress() bool {
	resultDaily, errDaily := uc.mSession.CountColl(uc.dailyColl)
	if errDaily != nil {
		// login
		uc.logging.WriteJsonLog("error",
			fmt.Sprintf("%s %s %s", "An error has occurred when trying to count", uc.dailyColl, "Collection"),
			uc.logging.GetTraceMsg())
	}
	resultAddress, errAddress := uc.mSession.CountColl("address")
	if errAddress != nil {
		// loggin
		uc.logging.WriteJsonLog("error",
			"An error has occurred when trying to count address Collection",
			uc.logging.GetTraceMsg())
	}
	if resultAddress > resultDaily {
		// loggin
		uc.logging.WriteJsonLog("info",
			"Address Collection is bigger than daily delivery Collection",
			uc.logging.GetTraceMsg())
		return false
	}
	return true
}

/**
 * Checks if the daily delivery collection exists
 */
func (uc *deliveryController) checkDeliveryCollection() bool {
	resultDaily, _ := uc.mSession.CountColl(uc.dailyColl)
	if resultDaily > 0 {
		uc.logging.WriteJsonLog("info",
			fmt.Sprintf("%s %s", uc.dailyColl, "Collection is has length bigger than 0"),
			uc.logging.GetTraceMsg())
		return true
	}
	uc.logging.WriteJsonLog("info",
		fmt.Sprintf("%s %s", uc.dailyColl, "Collection is empty"), uc.logging.GetTraceMsg())
	return false
}

/**
 * Creates a daily delivery collection
 */
func (uc *deliveryController) createDailyDeliveryCollection() {
	uc.logging.WriteJsonLog("info",
		fmt.Sprintf("%s %s %s", "We're going to ensure an index and the existence of the", uc.dailyColl, " Collection"),
		uc.logging.GetTraceMsg())
	deliveryDao := dao.NewDeliveryDao(uc.dailyColl)
	err := deliveryDao.CreateDailyCollection(uc.mSession, uc.dailyColl)
	if err != nil {
		uc.logging.WriteJsonLog("error",
			fmt.Sprintf("%s %s %s", "An error has occurred when trying to ensure", uc.dailyColl, "Collection"),
			uc.logging.GetTraceMsg())
	} else {
		uc.logging.WriteJsonLog("info",
			fmt.Sprintf("%s %s %s", "Collection", uc.dailyColl, "was successfully ensured"),
			uc.logging.GetTraceMsg())
	}
}

/**
 *
 */
func (uc *deliveryController) importDeliveryAddress() *models.Address {
	addressDao := dao.NewAddressDao()
	skip, _ := uc.mSession.CountColl(uc.dailyColl)
	uc.logging.WriteJsonLog("info",
		fmt.Sprintf("%s %s %s", "We're going to import an address, skipping", skip, "docs"),
		uc.logging.GetTraceMsg())
	address, _ := addressDao.GetAddress(uc.mSession, skip)
	uc.logging.WriteJsonLog("info",
		"We're going to update the daily delivery Collection docs",
		uc.logging.GetTraceMsg())
	return address
}

/**
 * Add the new address and
 */
func (uc *deliveryController) updateDeliveryAddresses() {
	uc.logging.WriteJsonLog("info",
		"We're going to update the daily delivery Collection docs",
		uc.logging.GetTraceMsg())
	address := uc.importDeliveryAddress()

	deliveryDao := dao.NewDeliveryDao(uc.dailyColl)
	delivery := models.Delivery{Address: address.Address, ZipCode: address.ZipCode, SussDlry: 0}
	deliveryDao.InsertDelivery(uc.mSession, &delivery)
}

/**
 *
 */
func (uc *deliveryController) getDailyDelivery() *models.Delivery {
	deliveryDao := dao.NewDeliveryDao(uc.dailyColl)
	delivery, _ := deliveryDao.GetDelivery(uc.mSession)
	delivery, _ = deliveryDao.IncrementField(uc.mSession, "sussDlry", delivery)
	return delivery
}

/**
 * @method GetDelivery
 */
func (uc *deliveryController) GetDelivery(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uc.dailyColl = time.Now().Format("2006_01_02")
	uc.logging = logit.NewSysLog()

	if uc.compareDeliveryAddress() {
		fmt.Println(uc.getDailyDelivery())
	} else if uc.checkDeliveryCollection() {
		uc.updateDeliveryAddresses()
		fmt.Println(uc.getDailyDelivery())
	} else {
		uc.createDailyDeliveryCollection()
		uc.updateDeliveryAddresses()
		fmt.Println(uc.getDailyDelivery())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
