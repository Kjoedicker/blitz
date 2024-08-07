package request

import (
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

	TargetNumber  int
	RequestGroup  int
	RequestNumber int
	ResponseTime  float64
	ErrorResponse error
}

type Requests []Request

func (requests Requests) CalculateAverageResponseTime() float64 {
	totalRequests := float64(len(requests))

	acc := 0.0
	for _, request := range requests {
		acc += request.ResponseTime
	}

	return acc / totalRequests
}

func (requests Requests) SumTotalErrors() int {
	totalErrors := 0
	for _, request := range requests {
		if request.ErrorResponse != nil {
			totalErrors++
		}
	}
	return totalErrors
}

func BuildCounter() func() int {
	var (
		totalRequests int
		mutex         sync.Mutex
	)

	return func() int {
		mutex.Lock()
		defer mutex.Unlock()

		totalRequests++
		return totalRequests
	}
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

func BuildTargets(plan plan.Plan) Requests {
	targets := Requests{}

	for targetNumber, target := range plan.Targets {

		target := Request{
			Hits:         target.Hits,
			Duration:     timing.IntToMinuteDuration(target.Duration),
			Interval:     timing.CalculateIntervalBetweenRequests(target.Hits, target.Interval),
			TargetNumber: targetNumber,
			Request: http.Request{
				Method: target.Method,
				URL:    buildUrl(plan.Host, target.Path),
				Header: target.Headers,
			},
		}

		targets = append(targets, target)
	}

	return targets
}
