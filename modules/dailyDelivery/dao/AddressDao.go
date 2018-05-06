package dao

import (
	"fmt"
	"github.com/jamesmallon/Alliggator"
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
 * @return *addressDao
 */
func NewAddressDao() *addressDao {
	return &addressDao{"address"}
}

/*
 * @method RunQuery
 */
func (us *addressDao) GetAddress(db *database.MongoSession, skip int) (*models.Address, error) {
	var wg sync.WaitGroup
	var address models.Address
	addrChannel := make(chan models.TrcReturn) // creates a new channel

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		jsonStr := fmt.Sprintf("%s%d%s", `[{"$project": {"_id": 0}},{"$skip": `, skip, `},{"$limit": 1}]`)
		addrChannel <- models.TrcReturn{
			Err:    db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&address),
			Result: address,
		}
		defer wg.Done()
	}()
	resp := <-addrChannel
	defer close(addrChannel)
	wg.Wait()
	return resp.Result.(*models.Address), resp.Err
}
