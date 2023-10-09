package main

import (
	"log"
	"math"
	"time"
)

func timer(identifier int) func() {
	start := time.Now()
	return func() {
		seconds := time.Since(start).Seconds()
		log.Printf("Call %d took: %fs \n", identifier, seconds)
	}
}

func HitsPer(totalRequests int, interval int) (HitsPer int) {
	return int(math.Round(float64(totalRequests) / float64(interval)))
}

func MinutesToSeconds(minutes int) int {
	return minutes * 60
}
