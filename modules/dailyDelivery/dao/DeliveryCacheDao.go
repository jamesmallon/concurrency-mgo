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

func (us *deliveryCacheDao) GettingKey(key string) string {
	redis := cache.ConnRedis()
	val := redis.Get(key)
	return val
}

func (us *deliveryCacheDao) SettingKey(key string) bool {
	redis := cache.ConnRedis()
	redis.Set(key, "1", 10000)
	return true
}
