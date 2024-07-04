package main

import (
	"sync"
	"time"

	"github.com/Kjoedicker/blitz/request"
	"github.com/Kjoedicker/blitz/timing"
)

func executeSynchronizedRequests(requestPrototype request.Request) {

	queue := sync.WaitGroup{}
	ticker := time.NewTicker(requestPrototype.Interval)
	stopMakingRequests := make(chan bool)

	requests := make(request.Requests, requestPrototype.Hits)
	returnRequestNumber := request.BuildCounter()

	for {
		select {
		case <-stopMakingRequests:
			queue.Wait()
			go requests.PrintResults()
			return
		case <-ticker.C:
			queue.Add(1)

			go func() {
				defer queue.Done()

				requestClone := requestPrototype
				requestNumber := returnRequestNumber()

				maxRequestsForIntervalReached := requestNumber >= requestPrototype.Hits+1
				if maxRequestsForIntervalReached {
					stopMakingRequests <- true
					return
				}

				request.Call(&requestClone)

				// This stores the request context at the index of when it was called.
				// This is useful for later printing the details of the request.
				requests[requestNumber-1] = requestClone
			}()
		}
	}
}

func Execute(requestPrototype request.Request) {

	done := make(timing.TimedChannel)
	go done.After(requestPrototype.Duration)

	for {
		select {
		case <-done:
			return
		default:
			go executeSynchronizedRequests(requestPrototype)

			// This helps compensate for the time it takes to allocate related resources
			// such as spinning up go routines and incrementing `requestGroup`.
			// Otherwise, overtime this adds up, and equates to missing a group of requests.
			time.Sleep(time.Second - requestPrototype.Interval)
		}
	}
}
