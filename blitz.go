package main

import (
	"sync"
	"time"
)

func trackRequestCount() func() int {
	totalRequests := 0
	return func() int {
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
			request := requestStructure
			request.Number = returnRequestNumber()

			request.Call()

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
