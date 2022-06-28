package test

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000)
	log.Printf("%v", n)
}
