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
func (us *addressDao) GetAddress(wg *sync.WaitGroup, db *database.MongoSession) (*models.Address, error) {
	var address models.Address
	wg.Add(1)
	c := make(chan *models.Address) // creates a new channel

	go func() {
		defer db.GetSession().Close()

		jsonStr := `[{"$sort": {"_id": 1}},{"$limit": 1}]`
		err := db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&address)
		if err != nil {
			fmt.Println("GetAddress ERROR:", err)
			return
		}
		c <- &address
	}()
	defer wg.Done()
	return <-c, nil
}
