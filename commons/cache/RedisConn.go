package cache

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var once sync.Once

type RedisClient struct {
	client *redis.Client
}

/**
 *
 */
func ConnRedis() *RedisClient {
	return &RedisClient{}
}

type response struct {
	result interface{}
	err    error
}

func (ch *RedisClient) connect() *redis.Client {
	once.Do(func() {
		ch.client = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			//PoolSize:           10,
			//PoolTimeout:        3 * time.Second,
			//DialTimeout:        10 * time.Second,
			//ReadTimeout:        30 * time.Second,
			//WriteTimeout:       30 * time.Second,
			//IdleTimeout:        500 * time.Millisecond,
			//IdleCheckFrequency: 500 * time.Millisecond,
			//Password: "", // no password set
			//DB:       0,  // use default DB
		})

	})
	return ch.client
}

func (ch *RedisClient) GetKey(key string) (string, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	respChannel := make(chan response)
	// singleton is thread safe and could be used with goroutines
	go func() {
		res, err := ch.connect().Get(key).Result()
		respChannel <- response{
			err:    err,
			result: res,
		}
		defer wg.Done()
	}()
	res := <-respChannel
	defer close(respChannel)
	wg.Wait()
	return res.result.(string), res.err
}

func (ch *RedisClient) SetKey(key string, val string) error {
	var wg sync.WaitGroup
	errChannel := make(chan error) // creates a new channel
	wg.Add(1)
	// singleton is thread safe and could be used with goroutines
	go func() {
		errChannel <- ch.connect().Set(key, val, 0).Err()
		defer wg.Done()
	}()
	err := <-errChannel
	defer close(errChannel)
	wg.Wait()
	return err
}

/**
 * @method SetTemporaryKey Sets the temporary key
 */
func (ch *RedisClient) SetTemporaryKey(key string, val string, milli int) error {
	var wg sync.WaitGroup
	errChannel := make(chan error) // creates a new channel
	wg.Add(1)
	// singleton is thread safe and could be used with goroutines
	go func() {
		ch.SetKey(key, val)
		errChannel <- ch.connect().PExpire(key, time.Duration(milli)*time.Millisecond).Err()
		defer wg.Done()
	}()
	err := <-errChannel
	defer close(errChannel)
	wg.Wait()
	return err
}

/**
 * @method IncrementKey Increments a key
 */
func (ch *RedisClient) IncrementKey(key string) (int64, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	respChannel := make(chan response)
	// singleton is thread safe and could be used with goroutines
	go func() {
		result, err := ch.connect().Incr(key).Result()
		respChannel <- response{
			result: result,
			err:    err,
		}
		defer wg.Done()
	}()
	res := <-respChannel
	defer close(respChannel)
	wg.Wait()
	return res.result.(int64), res.err
}
