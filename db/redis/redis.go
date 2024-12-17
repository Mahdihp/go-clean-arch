package redis

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

type RedisDb struct {
	client *redis.Client
}

func NewRedis(config config.RedisMarket) RedisDb {
	dbName, _ := strconv.Atoi(config.DBName)
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.HttpPort),
		Password:     config.Password,
		DB:           dbName,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return RedisDb{client: rdb}
}

func (a RedisDb) Client() *redis.Client {
	return a.client
}
