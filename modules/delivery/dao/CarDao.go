package dao

import (
	"fmt"
	"github.com/johnthegreenobrien/Alliggator"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"userv/commons/database"
	"userv/modules/delivery/models"
)

/**
 *
 */
type carDao struct {
	coll string
}

/**
 * @return *CarDao
 */
func NewCarDao() *carDao {
	return &carDao{"ipPorts"}
}

func (us *carDao) RunQuery(waitGroup *sync.WaitGroup) {
	db := database.ConnMongo()

	defer waitGroup.Done()

	sessionCopy := db.GetSession().Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB("delivery").C(us.coll)

	// Retrieve the list of stations.
	var deliveries models.Delivery

	jsonStr := `[{"$sort": {"counter": 1, "_id": 1}},{"$limit": 1}]`
	err := collection.Pipe(alliggator.Piperize(jsonStr)).One(&deliveries)
	if err != nil {
		fmt.Printf("RunQuery : ERROR : %s\n", err)
		return
	}

	change := db.GetIncrementer("counter")
	_, err = collection.Find(bson.M{"ipPort": deliveries.IpPort}).Apply(change, &deliveries)
	if err != nil {
		return
	}

	fmt.Println("Delivery:", deliveries)
}

//func (us *carDao) IncDomainCount(domain string, ipPort string) bool {
//	db := database.ConnMongo()
//	change := db.GetIncrementer("counter")
//
//	domainModel := models.Domain{}
//	_, err := db.GetCollection(domain).Find(bson.M{"ipPort": ipPort}).Apply(change, &domainModel)
//	if err != nil {
//		return false
//	}
//	return true
//}
