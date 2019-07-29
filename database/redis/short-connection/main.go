package main

import (
	"time"

	"gopkg.in/redis.v6"
)

// https://godoc.org/gopkg.in/redis.v6

//------------- Client ------------------------------------
var _client *redis.ClusterClient

//GetRedisClient 初始化redis连接并返回client
func GetRedisClient() *redis.ClusterClient {
	if _client != nil {
		return _client
	}

	_client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: "conf.Config.Redis.Host",
	})
	// log.Debugf("Connect to redis %v", conf.Config.Redis)

	_, err := _client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return _client
}

//------------- Redis Lock ------------------------------------
func RedisLock(key string, expire time.Duration) (bool, error) {
	ok, err := GetRedisClient().SetNX(key, key, expire).Result()
	// log.Debugf("redis lock %v", key)
	return ok, err
}

func RedisUnLock(key string) {
	GetRedisClient().Del(key)
	// log.Debugf("redis unlock %v", key)
	return
}

//-------------- Redis HTML Server ----------------------------
func RedisSetStaticHtml(key, html string) (err error) {
	return
}

func RedistGetStaticHTML(key string) (html string, err error) {
	return
}

func main() {
	//
}
