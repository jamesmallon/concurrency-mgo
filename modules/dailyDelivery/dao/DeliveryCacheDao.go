package dao

import (
	//	"fmt"
	//	"strconv"
	//"sync"
	"userv/commons/cache"
)

/**
 * @method deliveryCacheDao
 */
type deliveryCacheDao struct {
}

/**
 * @return *DeliveryCacheDao
 */
func NewDeliveryCacheDao() *deliveryCacheDao {
	return &deliveryCacheDao{}
}

func (us *deliveryCacheDao) ReadingKey(key string) string {
	//var waitGroup sync.WaitGroup
	redis := cache.ConnRedis()
	//val := redis.Get(key, &waitGroup)
	val := redis.Get(key)
	//waitGroup.Wait()
	return val
}

func (us *deliveryCacheDao) SettingKey(key string) bool {
	//var waitGroup sync.WaitGroup
	redis := cache.ConnRedis()
	//redis.Set(key, "1", 10000, &waitGroup)
	redis.Set(key, "1", 10000)
	return true
}
