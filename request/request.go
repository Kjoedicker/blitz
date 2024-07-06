package request

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/Kjoedicker/blitz/plan"
	"github.com/Kjoedicker/blitz/timing"
)

type Request struct {
	Request http.Request

	Hits     int
	Duration time.Duration
	Interval time.Duration

	RequestGroup  int
	RequestNumber int
	ResponseTime  float64
	ErrorResponse error
}

type Requests []Request

func BuildCounter() func() int {
	var totalRequests int
	var mutex sync.Mutex

	return func() int {
		mutex.Lock()
		defer mutex.Unlock()

		totalRequests++
		return totalRequests
	}
}

func (requests Requests) PrintResults() {
	for _, request := range requests {
		request.PrintResult()
	}
}

func (request Request) PrintResult(requestNumber int) {
	fmt.Printf("Request %d: %f seconds \n", requestNumber, request.ResponseTime)
}

// This is shared so connections can
// be reused thanks to internal caching
var client = http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     30 * time.Second,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 100,
	},
}

func Call(request *Request) (*http.Response, error) {

	stop := timing.BuildTimer()
	res, err := client.Do(&request.Request)
	request.ResponseTime = stop()

	if err != nil {
		request.ErrorResponse = err
		return nil, err
	}

	// The close is done after `err` is evaluated on purpose.
	// There is nothing left to close if an error occured.
	defer res.Body.Close()

	return res, nil
}

func buildUrl(host string, path string) *url.URL {
	rawUrl, err := url.JoinPath(host, path)
	if err != nil {
		log.Fatal(err)
	}
	url, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatal(err)
	}
	return url
}

func BuildRequests(plan plan.Plan) map[int]Request {
	targetRequests := make(map[int]Request)

	for targetNumber, target := range plan.Targets {

		targetRequests[targetNumber] = Request{
			Hits:     target.Hits,
			Duration: timing.IntToMinuteDuration(target.Duration),
			Interval: timing.CalculateIntervalBetweenRequests(target.Hits, target.Interval),
			Request: http.Request{
				Method: target.Method,
				URL:    buildUrl(plan.Host, target.Path),
				Header: target.Headers,
			},
		}
	}

	return targetRequests
}
