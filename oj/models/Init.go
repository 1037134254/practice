package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm/mysql"
)

var DB = Init()

var RDB = InitRedisDB()

func Init() *gorm.DB {
	args := "root:root1234@(127.0.0.1:13306)/OnlinePractice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
func InitRedisDB() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	return rdb
}
