package dao

import (
	"fmt"
	"microservices/api/modules/flowcontrol/models"
	"userv/commons/database"

	"gopkg.in/mgo.v2/bson"
)

/**
 *
 */
type CarDao struct {
	coll string
}

/**
 * @return *CarDao
 */
func DomainDAO() *CarDao {
	return &CarDao{"domain"}
}

/**
 * @param string jsonStr
 * @return interface{}
 */
func (us CarDao) CheckDomainBkp(jsonStr string) interface{} {
	db := database.ConnMongo()
	result := []models.Domain{}
	err := db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
