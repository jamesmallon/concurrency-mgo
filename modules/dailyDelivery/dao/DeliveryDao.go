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
 * @method deliveryDao
 */
type deliveryDao struct {
	coll string
}

/**
 * @return *DeliveryDao
 */
func NewDeliveryDao() *deliveryDao {
	return &deliveryDao{"delivery"}
}

/*
 * @method GetDelivery
 */
func (us *deliveryDao) GetDelivery(wg *sync.WaitGroup, db *database.MongoSession) (*models.Delivery, error) {
	c := make(chan *models.Delivery) // creates a new channel
	var delivery *models.Delivery

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		jsonStr := `[{"$sort": {"sussDlry": 1,"_id": 1}},{"$limit": 1}]`
		err := db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&delivery)
		if err != nil {
			fmt.Println("GetDelivery ERROR:", err)
			return
		}
		c <- delivery
		defer wg.Done()
	}()
	delivery = <-c
	defer close(c)
	return delivery, nil
}

/**
 * @method IncrementField
 */
func (us *deliveryDao) IncrementField(wg *sync.WaitGroup, db *database.MongoSession, field string, delivery *models.Delivery) (*models.Delivery, error) {
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
		defer wg.Done()
	}()
	delivery = <-c
	defer close(c)
	return delivery, nil
}

/**
 * @method CreateDailyCollection
 */
func (us *deliveryDao) CreateDailyCollection(wg *sync.WaitGroup, db *database.MongoSession, collName string) {
	wg.Add(1)
	go func() {
		// create unique index for zip-code field
		index := db.GetIndexObj([]string{"zipCode"}, true, false, false, false)

		err := db.GetCollection(collName).EnsureIndex(index)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()
}
