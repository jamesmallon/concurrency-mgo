package cache

import (
	"gopkg.in/go-redis/cache.v5"
	"gopkg.in/redis.v5"
	"sync"
)

var codec *cache.Codec
var once sync.Once

func GetInstance() *cache.Codec {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
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

func Get(keys []string, wg *sync.WaitGroup) chan SomeObj {
	wanted_objs := make(chan *SomeObj)
	for i, k := range keys {
		wg.Add(1)
		// singleton is thread safe and could be used with goroutines
		go func() {
			codec := GetInstance()
			wanted := new(SomeObj)
			if err := codec.Get(key, wanted); err == nil {
				wanted_objs <- wanted
			}
			wg.Done()
		}()
	}
	return wanted_objs
}

func Set(keys []string, vals []SomeObj, wg *sync.WaitGroup) {
	for i, k := range keys {
		wg.Add(1)
		// singleton is thread safe and could be used with goroutines
		go func() {
			codec := GetInstance()

			codec.Set(&cache.Item{
				Key:        k,
				Object:     vals[i],
				Expiration: time.Hour,
			})
			wg.Done()
		}()
	}
}
