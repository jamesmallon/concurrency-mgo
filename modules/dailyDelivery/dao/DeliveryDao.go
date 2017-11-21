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
 *
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

func (us *deliveryDao) RunQuery(waitGroup *sync.WaitGroup, db *database.MongoSession) {

	defer waitGroup.Done()

	sessionCopy := db.GetSession().Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB("delivery").C(us.coll)

	// Retrieve the list of stations.
	var deliveries models.Delivery

	jsonStr := `[{"$sort": {"sussDlry": 1, "_id": 1}},{"$limit": 1}]`
	err := collection.Pipe(alliggator.Piperize(jsonStr)).One(&deliveries)
	if err != nil {
		fmt.Printf("RunQuery : ERROR : %s\n", err)
		return
	}

	change := db.GetIncrementer("sussDlry")
	_, err = collection.Find(bson.M{"_id": deliveries.Id}).Apply(change, &deliveries)
	if err != nil {
		return
	}

	fmt.Println("Delivery:", deliveries)
}
