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
func NewDeliveryDao(collName string) *deliveryDao {
	return &deliveryDao{collName}
}

type response struct {
	delivery *models.Delivery
	err      error
}

func (us *deliveryDao) InsertDelivery(db *database.MongoSession, delivery *models.Delivery) error {
	var wg sync.WaitGroup
	errChannel := make(chan error) // creates a new channel
	var err error
	wg.Add(1)
	go func() {
		errChannel <- db.GetCollection(us.coll).Insert(delivery)
		if err != nil {
			fmt.Println(err)
		}
		defer wg.Done()
	}()
	err = <-errChannel
	defer close(errChannel)
	wg.Wait()
	return err
}

/*
 * @method GetDelivery
 */
func (us *deliveryDao) GetDelivery(db *database.MongoSession) (*models.Delivery, error) {
	var wg sync.WaitGroup
	dlrChannel := make(chan response)
	var dlry *models.Delivery

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		jsonStr := `[{"$sort": {"sussDlry": 1,"_id": 1}},{"$limit": 1}]`
		dlrChannel <- response{
			err:      db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&dlry),
			delivery: dlry,
		}
		defer wg.Done()
	}()
	resp := <-dlrChannel
	defer close(dlrChannel)
	wg.Wait()
	return resp.delivery, resp.err
}

/**
 * @method IncrementField
 */
func (us *deliveryDao) IncrementField(db *database.MongoSession, field string, deliveryPointer *models.Delivery) (*models.Delivery, error) {
	var wg sync.WaitGroup
	dlrChannel := make(chan response)
	var dlry *models.Delivery

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		change := db.GetIncrementer(field)
		_, errFd := db.GetCollection(us.coll).Find(bson.M{"_id": deliveryPointer.Id}).Apply(change, &dlry)
		dlrChannel <- response{
			err:      errFd,
			delivery: dlry,
		}
		defer wg.Done()
	}()
	resp := <-dlrChannel
	defer close(dlrChannel)
	wg.Wait()
	return resp.delivery, resp.err
}

/**
 * @method CreateDailyCollection
 */
func (us *deliveryDao) CreateDailyCollection(db *database.MongoSession, collName string) error {
	var wg sync.WaitGroup
	var err error
	errChannel := make(chan error) // creates a new channel
	wg.Add(1)
	go func() {
		// create unique index for zip-code field
		index := db.GetIndexObj([]string{"zipCode"}, true, false, false, false)

		errChannel <- db.GetCollection(collName).EnsureIndex(index)
		if err != nil {
			fmt.Println("CreateDailyCollections ERROR:", err)
		}
		defer wg.Done()
	}()
	err = <-errChannel
	defer close(errChannel)
	wg.Wait()
	return err
}
