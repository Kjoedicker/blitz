package main

import (
	"errors"
	"time"
)

func timer() func() float64 {
	start := time.Now()
	return func() float64 {
		seconds := time.Since(start).Seconds()
		return seconds
	}
}

// Calculates the interval needed to achieve a certain
// number of requests per unit of time (second or minute).
func CalculateIntervalBetweenRequests(requestsPerInterval int, unit string) (time.Duration, error) {
	switch unit {
	case "second":
		return time.Second / time.Duration(requestsPerInterval), nil
	case "minute":
		return time.Minute / time.Duration(requestsPerInterval), nil
	default:
		return 0, errors.New("invalid unit, must be 'second' or 'minute'")
	}
}

func MinutesToSeconds(minutes int) int {
	return minutes * 60
}

type TimedChannel chan bool

func (timeChannel TimedChannel) After(duration time.Duration) {
	time.Sleep(duration)
	timeChannel <- true
}
