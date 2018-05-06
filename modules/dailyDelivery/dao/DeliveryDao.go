package dao

import (
	"github.com/jamesmallon/Alliggator"
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

func (us *deliveryDao) InsertDelivery(db *database.MongoSession, delivery *models.Delivery) error {
	var wg sync.WaitGroup
	errChannel := make(chan error) // creates a new channel
	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		errChannel <- db.GetCollection(us.coll).Insert(delivery)
		defer wg.Done()
	}()
	err := <-errChannel
	defer close(errChannel)
	wg.Wait()
	return err
}

/*
 * @method GetDelivery
 */
func (us *deliveryDao) GetDelivery(db *database.MongoSession) (*models.Delivery, error) {
	var wg sync.WaitGroup
	dlrChannel := make(chan models.TrcReturn)
	var dlry *models.Delivery

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		jsonStr := `[{"$sort": {"sussDlry": 1,"_id": 1}},{"$limit": 1}]`
		dlrChannel <- models.TrcReturn{
			Err:    db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&dlry),
			Result: dlry,
		}
		defer wg.Done()
	}()
	resp := <-dlrChannel
	defer close(dlrChannel)
	wg.Wait()
	return resp.Result.(*models.Delivery), resp.Err
}

/**
 * @method IncrementField
 */
func (us *deliveryDao) IncrementField(db *database.MongoSession, field string, deliveryPointer *models.Delivery) (*models.Delivery, error) {
	var wg sync.WaitGroup
	dlrChannel := make(chan models.TrcReturn)
	var dlry *models.Delivery

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		change := db.GetIncrementer(field)
		_, errFd := db.GetCollection(us.coll).Find(bson.M{"_id": deliveryPointer.Id}).Apply(change, &dlry)
		dlrChannel <- models.TrcReturn{
			Err:    errFd,
			Result: dlry,
		}
		defer wg.Done()
	}()
	resp := <-dlrChannel
	defer close(dlrChannel)
	wg.Wait()
	return resp.Result.(*models.Delivery), resp.Err
}

/**
 * @method CreateDailyCollection
 */
func (us *deliveryDao) CreateDailyCollection(db *database.MongoSession, collName string) error {
	var wg sync.WaitGroup
	errChannel := make(chan error) // creates a new channel
	wg.Add(1)
	go func() {
		defer db.GetSession().Close()
		// create unique index for zip-code field
		index := db.GetIndexObj([]string{"zipCode"}, true, false, false, false)
		errChannel <- db.GetCollection(collName).EnsureIndex(index)
		defer wg.Done()
	}()
	err := <-errChannel
	defer close(errChannel)
	wg.Wait()
	return err
}
