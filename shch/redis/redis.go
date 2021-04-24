package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

func main() {
	//连接服务器
	redisClient:= redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	//心跳
	pong, err := redisClient.Ping().Result()
	log.Println(pong, err) // Output: PONG <nil>

	//kv读写
	err = redisClient.Set("key", "value", 1*time.Second).Err()
	log.Println(err)

	//获取过期时间
	tm, err := redisClient.TTL("key").Result()
	log.Println(tm)

	val, err := redisClient.Get("key").Result()
	log.Println(val, err)

	val2, err := redisClient.Get("missing_key").Result()
	if err == redis.Nil {
		log.Println("missing_key does not exist")
	} else if err != nil {
		log.Println("missing_key", val2, err)
	}

	//不存在才设置 过期时间 nx ex
	value, err := redisClient.SetNX("counter", 0, 1*time.Second).Result()
	log.Println("setnx", value, err)

	//Incr
	result, err := redisClient.Incr("counter").Result()
	log.Println("Incr", result, err)
}
