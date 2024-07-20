package main

import (
	"sync"
	"time"

	"github.com/Kjoedicker/blitz/request"
	"github.com/Kjoedicker/blitz/results"
	"github.com/Kjoedicker/blitz/timing"
)

func trigger(target request.Request) {

	queue := sync.WaitGroup{}
	ticker := time.NewTicker(target.Interval)
	stopMakingRequests := make(chan bool)

	requests := make(request.Requests, target.Hits)
	returnRequestNumber := request.BuildCounter()

	for range ticker.C {
		select {
		case <-stopMakingRequests:
			queue.Wait()
			results.PrintAll(requests)
			return
		default:
			queue.Add(1)

			go func() {
				defer queue.Done()

				requestClone := target
				requestClone.RequestNumber = returnRequestNumber()

				maxRequestsForIntervalReached := requestClone.RequestNumber >= target.Hits+1
				if maxRequestsForIntervalReached {
					stopMakingRequests <- true
					return
				}

				request.Call(&requestClone)

				// This stores the request context at the index of when it was called.
				// This is useful for later printing the details of the request.
				requests[requestClone.RequestNumber-1] = requestClone
			}()
		}
	}
}

func Target(target request.Request) {

	returnGroupNumber := request.BuildCounter()

	done := make(timing.TimedChannel)
	go done.After(target.Duration)

	var wg = sync.WaitGroup{}
	for {
		select {
		case <-done:
			wg.Wait()
			return
		default:
			target.RequestGroup = returnGroupNumber()

			wg.Add(1)
			go func() {
				defer wg.Done()
				trigger(target)
			}()

			stop := timing.BuildTimer()
			wg.Wait()

			var (
				waitTime   = time.Duration(stop())
				blockTime  = time.Duration(time.Second - target.Interval)
				totalSleep = blockTime - waitTime
			)

			// This helps compensate for the time it takes to allocate related resources
			// such as spinning up go routines and incrementing `requestGroup`.
			// Otherwise, overtime this adds up, and equates to missing a group of requests.
			time.Sleep(totalSleep)
		}
	}
}

func Setup(targets request.Requests) {
	for _, target := range targets {
		Target(target)
	}
}
