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
type deliveryCacheDao struct{}

/**
 * @return *DeliveryCacheDao
 */
func NewDeliveryCacheDao() *deliveryCacheDao {
	return &deliveryCacheDao{}
}

func (us *deliveryCacheDao) GettingKey(key string, rClient *cache.RedisClient) string {
	val, err := rClient.GetKey(key)
	if err != nil {
	}
	return val
}

func (us *deliveryCacheDao) SettingKey(key string, val string, rClient *cache.RedisClient) bool {
	err := rClient.SetKey(key, val)
	if err != nil {
	}
	return true
}

func (us *deliveryCacheDao) SettingTempKey(key string, val string, rClient *cache.RedisClient) bool {
	err := rClient.SetTemporaryKey(key, val, 10000)
	if err != nil {
	}
	return true
}

func (us *deliveryCacheDao) IncrementingKey(key string, rClient *cache.RedisClient) bool {
	_, err := rClient.IncrementKey(key)
	if err != nil {
	}
	return true
}
