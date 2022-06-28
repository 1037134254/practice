package test

import (
	"fmt"
	"github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"strconv"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Printf(strconv.Itoa(time.Now().Year()))
	month := time.Now().Month()
	fmt.Printf(strconv.Itoa(int(month)))
	day := time.Now().Day()
	fmt.Printf(strconv.Itoa(int(day)))
}

func TestDss(t *testing.T) {
	rate, _ := limiter.NewRateFromFormatted("1024-S")
	store, _ := sredis.NewStore(rdb)
	l := limiter.New(store, rate)
	fmt.Println(l)
}
