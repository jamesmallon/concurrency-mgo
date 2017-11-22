package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/johnthegreenobrien/Alliggator"
	"userv/modules/dailyDelivery/models"
	//"reflect"
	"sync"
	"time"
)

/**
 * This file implements the singleton pattern to avoid the sistem to deal with more
 * than one Mongodb connection at once.
 */
const (
	MongoDBHosts = "localhost:27307"
	AuthDatabase = "delivery"
	AuthUserName = "deliveryUser"
	AuthPassword = "delivery123"
	Database     = "delivery"
)

/**
 *
 */
type MongoSession struct {
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection
}

var mgoSession *mgo.Session
var once sync.Once

func GetSession() *mgo.Session {
	once.Do(func() {
		var err error
		mgoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{MongoDBHosts},
			//Timeout: 60 * time.Second,
			Timeout:  1 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		})
		if err != nil {
			fmt.Println("CreateSession: %s\n", err)
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	})
	return mgoSession
}

func GetThat(wg *sync.WaitGroup) (*models.Address, error) {
	var address models.Address
	c := make(chan *models.Address) // creates a new channel

	wg.Add(1)
	go func() {
		session := GetSession().Copy()
		defer session.Close()

		jsonStr := `[{"$sort": {"_id": 1}},{"$limit": 1}]`
		err := session.DB(Database).C("delivery").Pipe(alliggator.Piperize(jsonStr)).One(&address)
		if err != nil {
			fmt.Println("GetAddress ERROR:", err)
			return
		} else {
			fmt.Println("Mongo is ongoing")
		}
		c <- &address
		wg.Done()
	}()

	return <-c, nil
}

//func RunQuery(waitGroup *sync.WaitGroup, mongoSession *mgo.Session) *models.Delivery {
//	defer waitGroup.Done()
//
//	sessionCopy := mongoSession.Copy()
//	defer sessionCopy.Close()
//
//	// Get a collection to execute the query against.
//	collection := GetInstance.DB(Database).C("delivery")
//
//	// Retrieve the list of stations.
//	var delivery models.Delivery
//	err := collection.Find(nil).One(&delivery)
//	if err != nil {
//		fmt.Println("RunQuery ERROR:", err)
//		return delivery
//	}
//	return delivery
//}
//
//func GetIncrementer(field string) mgo.Change {
//	change := mgo.Change{
//		Update:    bson.M{"$inc": bson.M{field: 1}},
//		ReturnNew: true,
//	}
//	return change
//}
//
//func GetIndexObj(indexField []string,
//	unique bool,
//	dropDups bool,
//	background bool,
//	sparse bool) mgo.Index {
//
//	index := mgo.Index{
//		Key:        indexField,
//		Unique:     unique,
//		DropDups:   dropDups,
//		Background: background, // See notes.
//		Sparse:     sparse,
//	}
//	return index
//}
