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
	val := rClient.GetKey(key)
	return val
}

func (us *deliveryCacheDao) SettingKey(key string, val string, rClient *cache.RedisClient) bool {
	rClient.SetKey(key, val)
	return true
}

func (us *deliveryCacheDao) SettingTempKey(key string, val string, rClient *cache.RedisClient) bool {
	rClient.SetTemporaryKey(key, val, 10000)
	return true
}

func (us *deliveryCacheDao) IncrementingKey(key string, rClient *cache.RedisClient) bool {
	rClient.IncrementKey(key)
	return true
}
