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

	Number int
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

func (request Request) Call() (*http.Response, error) {
	stop := timer(request.Number)
	res, err := client.Do(&request.Request)
	stop()

	if err != nil {
		fmt.Println(res.Status, res.StatusCode)
		return nil, err
	}

	fmt.Println(res.StatusCode)

	defer res.Body.Close()
	return res, nil
}
