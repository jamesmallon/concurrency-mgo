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

func (us *deliveryCacheDao) GettingKey(key string, rClient *cache.RedisClient) string {
	val := rClient.Get(key)
	return val
}

func (us *deliveryCacheDao) SettingKey(key string, val string, rClient *cache.RedisClient) bool {
	rClient.Set(key, val, 10000)
	return true
}
