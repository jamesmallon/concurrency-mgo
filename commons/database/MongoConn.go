/**
 *
 */
package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

/**
 */
func ConnMongo() *MongoSession {
	return &MongoSession{}
}

var once sync.Once

func (db *MongoSession) connect() *MongoSession {
	once.Do(func() {
		var err error
		db.session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Timeout:  60 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		})
		if err != nil {
			fmt.Println("CreateSession: %s\n", err)
		} else {
			fmt.Println("Session Created")
		}

		db.session.SetMode(mgo.Monotonic, true)
		db.database = db.session.DB(Database)
	})
	return db
}

func (db *MongoSession) GetSession() *mgo.Session {
	//return db.session.Copy()
	return db.connect().session.Copy()
}

/**
 * @method db.db UseDB
 * @param string db
 * @return *mgo.Database db.db
 */
func (db *MongoSession) UseDB(dbase string) *MongoSession {
	db.database = db.session.DB(dbase)
	return db
}

/**
 * @method db.collection GetCollection
 * @param string coll
 * @return *mgo.Collection db.collection
 */
func (db *MongoSession) GetCollection(coll string) *mgo.Collection {
	db.collection = db.database.C(coll)
	return db.collection
}

func (db *MongoSession) GetIncrementer(field string) mgo.Change {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{field: 1}},
		ReturnNew: true,
	}
	return change
}

func (db *MongoSession) GetIndexObj(indexField []string,
	unique bool,
	dropDups bool,
	background bool,
	sparse bool) mgo.Index {

	index := mgo.Index{
		Key:        indexField,
		Unique:     unique,
		DropDups:   dropDups,
		Background: background, // See notes.
		Sparse:     sparse,
	}
	return index
}

/**
 *
 */
type counter struct {
	Count int `json:"count" bson:"count,omitempty"`
}

type response struct {
	counter int
	err     error
}

/**
 *
 */
func (db *MongoSession) CountColl(collStr string) (int, error) {
	var wg sync.WaitGroup
	countChannel := make(chan response) // creates a new channel
	var result counter

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		countChannel <- response{
			err:     db.GetCollection(collStr).Pipe([]bson.M{bson.M{"$group": bson.M{"_id": "count", "count": bson.M{"$sum": 1}}}}).One(&result),
			counter: result.Count,
		}
		defer wg.Done()
	}()
	resp := <-countChannel
	defer close(countChannel)
	wg.Wait()
	return resp.counter, resp.err
}
