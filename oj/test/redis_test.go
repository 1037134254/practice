package test

import (
	"context"
	"example.com/m/v2/models"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "123456", // no password set
	DB:       0,        // use default DB
})

func TestSet(t *testing.T) {
	err := models.InitRedisDB().Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	get, _ := models.InitRedisDB().Get(ctx, "2865549101@qq.com").Result()
	log.Printf("key:%s", get)
}
