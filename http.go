package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Request struct {
	Method  string
	Url     url.URL
	rawUrl  string
	Headers http.Header
	Request http.Request

	RequestGroup int
	Number       int
	ResponseTime float64
}

type Requests []Request

func (requests Requests) PrintResults() {
	for index, request := range requests {
		if index == 0 {
			fmt.Println("\nRequest group:", request.RequestGroup)
		}
		request.PrintResult()
	}
}

func (request Request) PrintResult() {
	fmt.Printf("Request %d: %f seconds \n", request.Number, request.ResponseTime)
}

func (request Request) Init() Request {
	url, err := url.Parse(request.rawUrl)
	if err != nil {
		log.Fatal(err)
	}

	request.Request = http.Request{
		Method: request.Method,
		URL:    url,
		Header: request.Headers,
	}
	return request
}

// This is shared so connections can
// be reused thanks to internal caching
var client = http.Client{}

func (request *Request) Call() (*http.Response, error) {
	stop := timer()
	res, err := client.Do(&request.Request)
	request.ResponseTime = stop()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}
