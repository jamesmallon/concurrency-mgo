package dao

import (
	"fmt"
	"github.com/johnthegreenobrien/Alliggator"
	//"gopkg.in/mgo.v2"
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
 * @method RunQuery
 */
func (us *deliveryDao) GetDeliveryAddress(wg *sync.WaitGroup, db *database.MongoSession, c chan *models.Delivery) {

	defer wg.Done()

	sessionCopy := db.GetSession().Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB("delivery").C(us.coll)

	// Retrieve the list of stations.
	var delivery models.Delivery

	jsonStr := `[{"$sort": {"counter": 1, "_id": 1}},{"$limit": 1}]`
	err := collection.Pipe(alliggator.Piperize(jsonStr)).One(&delivery)
	if err != nil {
		fmt.Printf("RunQuery : ERROR : %s\n", err)
		return
	}
	fmt.Println("Delivery:", delivery)
	c <- &delivery
}

/**
 * @method IncrementField
 */
func (us *deliveryDao) IncrementField(wg *sync.WaitGroup, db *database.MongoSession, field string, delivery *models.Delivery) {

	defer wg.Done()

	sessionCopy := db.GetSession().Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB("delivery").C(us.coll)
	change := db.GetIncrementer(field)
	_, err := collection.Find(bson.M{"_id": delivery.Id}).Apply(change, &delivery)
	if err != nil {
		return
	}
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
//func (us *deliveryDao) CreateCollection(collectionName string) {
//
//}
