package configurator

import (
	"fmt"
	redis "github.com/go-redis/redis/v8"
)

func OpenRedisConnection(conf *AppConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb, nil
}
