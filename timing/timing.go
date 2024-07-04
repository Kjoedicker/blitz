package timing

import (
	"errors"
	"time"
)

type TimedChannel chan bool

func (timeChannel TimedChannel) After(duration time.Duration) {
	time.Sleep(duration)
	timeChannel <- true
}

func BuildTimer() func() float64 {
	start := time.Now()
	return func() float64 {
		seconds := time.Since(start).Seconds()
		return seconds
	}
}

// Calculates the interval needed to achieve a certain
// number of requests per unit of time (second or minute).
func CalculateIntervalBetweenRequests(requestsPerInterval int, unit string) time.Duration {
	switch unit {
	case "second":
		return time.Second / time.Duration(requestsPerInterval)
	case "minute":
		return time.Minute / time.Duration(requestsPerInterval)
	default:
		err := errors.New("invalid unit, must be 'second' or 'minute'")
		panic(err)
	}
}

func MinutesToSeconds(minutes int) int {
	return minutes * 60
}

func IntToMinuteDuration(duration int) time.Duration {
	return time.Duration(duration) * time.Minute
}
