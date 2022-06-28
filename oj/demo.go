package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
	"net/http"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "123456", // no password set
	DB:       0,        // use default DB
})

func main() {
	rate, err := limiter.NewRateFromFormatted("200-H")
	if err != nil {
		log.Fatal(err)
		return
	}
	store, err := sredis.NewStoreWithOptions(rdb, limiter.StoreOptions{Prefix: "limiter_gin_example"})
	if err != nil {
		log.Fatal(err)
		return
	}
	// 创建中间件
	middleware := mgin.NewMiddleware(limiter.New(store, rate))
	engine := gin.Default()
	engine.ForwardedByClientIP = true
	engine.Use(middleware) // 使用中间件
	engine.GET("/", index)
	err = engine.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func CopySpeed() {
	rate, _ := limiter.NewRateFromFormatted("1024-S")
	store, _ := sredis.NewStore(rdb)
	limiter.New(store, rate)
}

func index(c *gin.Context) {
	type message struct {
		Message string `json:"message"`
	}
	resp := message{Message: "ok"}
	c.JSON(http.StatusOK, resp)
}
