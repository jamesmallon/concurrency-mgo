package database

import (
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"log"
	"sync"
	"time"
	"userv/modules/dailyDelivery/models"
)

const (
	MongoDBHosts = "localhost:27307"
	AuthDatabase = "delivery"
	AuthUserName = "deliveryUser"
	AuthPassword = "delivery123"
	TestDatabase = "delivery"
)

var mongoSession *mgo.Session

// main is the entry point for the application.
func ConnMongo() *mgo.Session {

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)
	return mongoSession
}

func RunQuery(waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
	defer waitGroup.Done()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(TestDatabase).C("delivery")

	var delivery models.Delivery
	err := collection.Find(nil).One(&delivery)
	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return
	}
}
