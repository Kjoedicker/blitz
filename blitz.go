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

func makeSynchronizedRequests(requestStructure Request, hitsPerSecond int) {
	queue := sync.WaitGroup{}

	requests := make(Requests, hitsPerSecond)
	returnRequestNumber := trackRequestCount()

	for totalRequests := 1; totalRequests <= hitsPerSecond; totalRequests++ {
		queue.Add(1)

		go func() {
			defer queue.Done()

			request := requestStructure
			request.Number = returnRequestNumber()

			request.Call()

			// This stores the request context at the
			// index of when it was called.
			requests[request.Number-1] = request

			queue.Done()
		}()
	}

	queue.Wait()

	requests.PrintResults()

	time.Sleep(time.Second)
}

func MakeRequestsForDuration(requestStructure Request, duration int, hitsPerSecond int) {

	done := make(TimedChannel)
	go done.After(time.Duration(duration) * time.Minute)

	for requestGroup := 0; ; requestGroup++ {
		request := requestStructure
		request.RequestGroup = requestGroup

		select {
		case <-done:
			return
		default:
			makeSynchronizedRequests(request, hitsPerSecond)
		}
	}
}
