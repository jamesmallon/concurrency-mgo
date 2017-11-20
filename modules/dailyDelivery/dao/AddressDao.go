package dao

import (
	"fmt"
	"github.com/johnthegreenobrien/Alliggator"
	"sync"
	"userv/commons/database"
	"userv/modules/dailyDelivery/models"
)

/**
 * @method addressDao
 */
type addressDao struct {
	coll string
}

/**
 * @return *DeliveryDao
 */
func NewAddressDao() *addressDao {
	return &addressDao{"address"}
}

/*
 * @method RunQuery
 */
func (us *addressDao) GetAddress(wg *sync.WaitGroup, mgoSess *database.MongoSession) (*models.Address, error) {
	//db := mgoSess.UseDB("delivery")
	var address models.Address
	wg.Add(1)
	c := make(chan *models.Address) // creates a new channel

	go func() {
		//db := mgoSess.GetSession().Copy()
		db := mgoSess.GetSession()
		defer db.Close()

		// Get a collection to execute the query against.
		collection := db.DB("delivery").C(us.coll)

		jsonStr := `[{"$sort": {"_id": 1}},{"$limit": 1}]`
		err := collection.Pipe(alliggator.Piperize(jsonStr)).One(&address)
		//err := db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&address)
		if err != nil {
			fmt.Printf("RunQuery : ERROR : %s\n", err)
			return
		}
		c <- &address
	}()
	defer wg.Done()
	return <-c, nil
}

/**
 * @method IncrementField
 */
//func (us *addressDao) IncrementField(wg *sync.WaitGroup, db *database.MongoSession, field string, delivery *models.Address) {
//
//	defer wg.Done()
//
//	sessionCopy := db.GetSession().Copy()
//	defer sessionCopy.Close()
//
//	// Get a collection to execute the query against.
//	collection := sessionCopy.DB("delivery").C(us.coll)
//	change := db.GetIncrementer(field)
//	_, err := collection.Find(bson.M{"_id": delivery.Id}).Apply(change, &delivery)
//	if err != nil {
//		return
//	}
//}

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
//func (us *addressDao) CreateCollection(collectionName string) {
//
//}
