package cache

import (
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"

	"github.com/go-redis/cache"
	"sync"
	"time"
)

var codec *cache.Codec
var once sync.Once

type RedisClient struct{}

/**
 * @method db.sessione.Clone() GetMongoSession It create and instantiate a Mongodb connection
 * @return db.sessione.Clone()
 */
func ConnRedis() *RedisClient {
	return &RedisClient{}
}

func (ch *RedisClient) connect() *cache.Codec {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
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

		codec = &cache.Codec{
			Redis: client,

			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		}
	})
	return codec
}

func (ch *RedisClient) Get(key string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan string)
	// singleton is thread safe and could be used with goroutines
	go func() {
		codec := ch.connect()
		var wanted string
		if err := codec.Get(key, &wanted); err == nil {
			c <- wanted
		}
		wg.Done()
	}()
	res := <-c
	defer close(c)
	wg.Wait()
	return res
}

func (ch *RedisClient) Set(key string, val string, milli int) {
	var wg sync.WaitGroup
	wg.Add(1)
	// singleton is thread safe and could be used with goroutines
	go func() {
		codec := ch.connect()
		codec.Set(&cache.Item{
			Key:        key,
			Object:     val,
			Expiration: time.Duration(milli) * time.Millisecond,
		})
		wg.Done()
	}()
	wg.Wait()
}
