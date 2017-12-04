package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var once sync.Once

type RedisClient struct {
	client *redis.Client
}

/**
 * @method db.sessione.Clone() GetMongoSession It create and instantiate a Mongodb connection
 */
func ConnRedis() *RedisClient {
	return &RedisClient{}
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

func (ch *RedisClient) GetKey(key string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan string)
	// singleton is thread safe and could be used with goroutines
	go func() {
		result, err := ch.connect().Get(key).Result()
		if err != nil {
			fmt.Println("Error getting redis key")
		}
		c <- result
		defer wg.Done()
	}()
	res := <-c
	defer close(c)
	wg.Wait()
	return res
}

func (ch *RedisClient) SetKey(key string, val string) {
	var wg sync.WaitGroup
	wg.Add(1)
	// singleton is thread safe and could be used with goroutines
	go func() {
		err := ch.connect().Set(key, val, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func (ch *RedisClient) SetTemporaryKey(key string, val string, milli int) {
	var wg sync.WaitGroup
	wg.Add(1)
	// singleton is thread safe and could be used with goroutines
	go func() {
		ch.SetKey(key, val)
		err := ch.connect().PExpire(key, time.Duration(milli)*time.Millisecond).Err()
		if err != nil {
			fmt.Println(err)
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func (ch *RedisClient) IncrementKey(key string) int64 {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int64)
	// singleton is thread safe and could be used with goroutines
	go func() {
		result, err := ch.connect().Incr(key).Result()
		if err != nil {
			fmt.Println(err)
		}
		c <- result
		defer wg.Done()
	}()
	res := <-c
	defer close(c)
	wg.Wait()
	return res
}
