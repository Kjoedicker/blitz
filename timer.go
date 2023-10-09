package main

import (
	"log"
	"time"
)

func timer(identifier int) func() {
	start := time.Now()
	return func() {
		seconds := time.Since(start).Seconds()
		log.Printf("Call %d took: %fs \n", identifier, seconds)
	}
}
