package main

import (
	"math"
	"time"
)

func timer() func() float64 {
	start := time.Now()
	return func() float64 {
		seconds := time.Since(start).Seconds()
		return seconds
	}
}

func HitsPer(totalRequests int, interval int) (HitsPer int) {
	return int(math.Round(float64(totalRequests) / float64(interval)))
}

func MinutesToSeconds(minutes int) int {
	return minutes * 60
}

type TimedChannel chan bool

func (timeChannel TimedChannel) After(duration time.Duration) {
	time.Sleep(duration)
	timeChannel <- true
}
