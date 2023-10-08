package main

import (
	"fmt"
	"time"
)

func timer(identifier int) func() {
	start := time.Now()
	return func() {
		seconds := time.Since(start).Seconds()
		fmt.Printf("Call %d took: %fs \n", identifier, seconds)
	}
}
