package test

import (
	"math"
	"time"
)

type Clock interface {
	Now() time.Time      //纳秒
	Sleep(time.Duration) // 经过的时间 纳秒
}

type Mode int8

const (
	SingleChannel Mode = 1 + iota //Read or Write 使用
	DualChannel                   //Copy方法使用
)

type fileLimiter struct {
	sleepFor   time.Duration //需要睡眠的时间
	start      time.Time     //开始时间
	clock      Clock         //时间
	rate       int64         //限制每秒读写速度
	tokens     int64         //可用令牌数
	latestTick float64       //上一次生成的tick
	mode       Mode
}

// NewFileLimiter 文件读写限速
func NewFileLimiter(rate int64) *fileLimiter {
	clock := new(sleepTime) // rate clo
	return &fileLimiter{
		start:      clock.Now(),
		rate:       rate,
		clock:      clock,
		sleepFor:   0,
		latestTick: 0,
		mode:       SingleChannel,
	}
}

// 返回持续时间
func (f *fileLimiter) currentTick(now time.Time) float64 {
	return now.Sub(f.start).Seconds()
}

// 当前令牌数 = 上一次剩余的令牌数 + (本次取令牌的时刻-上一次取令牌的时刻)/放置令牌的时间间隔 * 每次放置的令牌数
func (f *fileLimiter) currentTokens(tick float64) {
	f.tokens += int64(math.Ceil((tick - f.latestTick) * float64(f.rate)))
	f.latestTick = tick
}

//调整应该睡眠的时间
func (f *fileLimiter) sleepTime(now time.Time, count int64) {
	if count <= 0 {
		return
	}
	f.currentTokens(f.currentTick(now))
	f.tokens -= count
	if f.tokens < 0 {
		millisecond := math.Abs(float64(f.tokens)) / float64(f.rate) * 1000
		f.sleepFor = time.Millisecond * time.Duration(millisecond)
	} else {
		f.sleepFor = 0
	}
}

func (f *fileLimiter) Wait(takeTokens int64) {
	f.sleepTime(f.clock.Now(), takeTokens/int64(f.mode))
	if f.sleepFor > 0 {
		f.clock.Sleep(f.sleepFor)
	}
}

func (f *fileLimiter) Reset() {
	if f.clock == nil {
		f.clock = new(sleepTime)
	}
	f.start = f.clock.Now()
	f.sleepFor = 0
	f.latestTick = 0
}
