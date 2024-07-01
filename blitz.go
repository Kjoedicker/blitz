package main

import (
	"sync"
	"time"
)

func trackRequestCount() func() int {
	var totalRequests int
	var mutex sync.Mutex

	return func() int {
		mutex.Lock()
		defer mutex.Unlock()

		totalRequests++
		return totalRequests
	}
}

func makeSynchronizedRequests(requestStructure Request, totalHits int, intervalBetweenRequests time.Duration) {
	queue := sync.WaitGroup{}

	requests := make(Requests, totalHits)
	returnRequestNumber := trackRequestCount()
	done := make(chan bool)
	ticker := time.NewTicker(intervalBetweenRequests)

	for {
		select {
		case <-done:
			queue.Wait()
			go requests.PrintResults()
			return
		case <-ticker.C:
			queue.Add(1)

			go func() {
				defer queue.Done()

				request := requestStructure
				request.Number = returnRequestNumber()

				if request.Number >= totalHits+1 {
					done <- true
					return
				}

				request.Call()

				// This stores the request context at the index of when it was called.
				// This is useful for later printing the details of the request.
				requests[request.Number-1] = request

			}()
		}
	}
}

func MakeRequestsForDuration(requestStructure Request, totalHits int, intervalBetweenRequests time.Duration, duration int) {
	done := make(TimedChannel)
	go done.After(time.Duration(duration) * time.Minute)

	requestGroup := 1
	for {
		select {
		case <-done:
			return
		default:
			request := requestStructure
			request.RequestGroup = requestGroup

			go makeSynchronizedRequests(request, totalHits, intervalBetweenRequests)

			// This helps compensate for the time it takes to allocate related resources
			// such as spinning up go routines and incrementing `requestGroup`.
			// Otherwise, overtime this adds up, and equates to missing a group of requests.
			time.Sleep(time.Second - intervalBetweenRequests)

			requestGroup++
		}
	}
}
