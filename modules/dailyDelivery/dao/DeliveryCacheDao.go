package dao

import (
	"fmt"
	"github.com/johnthegreenobrien/Alliggator"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"userv/commons/database"
	"userv/modules/dailyDelivery/models"
)

/**
 * @method deliveryCacheDao
 */
type deliveryCacheDao struct {
	coll string
}

/**
 * @return *DeliveryCacheDao
 */
func NewDeliveryCacheDao() *deliveryCacheDao {
	return &deliveryCacheDao{"delivery"}
}

/*
 * @method GetDelivery
 */
func (us *deliveryCacheDao) GetDelivery(wg *sync.WaitGroup, db *database.MongoSession) (*models.Delivery, error) {
	var delivery models.Delivery
	c := make(chan *models.Delivery) // creates a new channel

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		jsonStr := `[{"$sort": {"sussDlry": 1,"_id": 1}},{"$limit": 1}]`
		err := db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&delivery)
		if err != nil {
			fmt.Println("GetDelivery ERROR:", err)
			return
		}
		c <- &delivery
	}()
	defer wg.Done()

	return <-c, nil
}

/**
 * @method IncrementField
 */
func (us *deliveryCacheDao) IncrementField(wg *sync.WaitGroup, db *database.MongoSession, field string, delivery *models.Delivery) (*models.Delivery, error) {
	c := make(chan *models.Delivery) // creates a new channel

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		change := db.GetIncrementer(field)
		_, err := db.GetCollection(us.coll).Find(bson.M{"_id": delivery.Id}).Apply(change, &delivery)
		if err != nil {
			fmt.Println("IncrementField ERROR:", err)
			return
		}
		c <- delivery
	}()
	defer wg.Done()

	return <-c, nil
}

//func (us *domainDAO) CreateCollIndex(domain string, indexField []string) {
//	db := database.ConnMongo()
//
//	index := db.GetIndexObj(indexField, true, false, false, false)
//
//	err := db.GetCollection(domain).EnsureIndex(index)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func (us *deliveryCacheDao) CreateCollection(collectionName string) {
//
//}
