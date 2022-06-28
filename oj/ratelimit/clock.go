package test

import (
	"time"
)

type sleepTime struct{}

func (s sleepTime) Now() time.Time {
	return time.Now()
}

func (s sleepTime) Sleep(duration time.Duration) {
	time.Sleep(duration)
}
