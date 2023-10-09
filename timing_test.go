package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_HitsPer(t *testing.T) {
	expectations := map[string][]int{
		"1000 requests every 60 seconds":   {1000, 60, 17},
		"10000 requests every 60 seconds":  {10000, 60, 167},
		"10000 requests every 120 seconds": {10000, 120, 83},
	}

	for description, values := range expectations {
		requests := values[0]
		interval := values[1]
		expectation := values[2]
		t.Run(description, func(t *testing.T) {
			hitsPerSecond := HitsPer(requests, interval)
			assert.Equal(t, expectation, hitsPerSecond)
		})
	}
}
