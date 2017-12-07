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
	c := make(chan *models.Address) // creates a new channel

	wg.Add(1)
	go func() {
		defer db.GetSession().Close()

		jsonStr := fmt.Sprintf("%s%d%s", `[{"$project": {"_id": 0}},{"$skip": `, skip, `},{"$limit": 1}]`)
		err := db.GetCollection(us.coll).Pipe(alliggator.Piperize(jsonStr)).One(&address)
		if err != nil {
			fmt.Println("GetAddress ERROR:", err)
		}
		c <- &address
	}()
	defer wg.Done()

	return <-c, nil
}
