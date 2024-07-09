package main

import (
	"github.com/Kjoedicker/blitz/cli"
	"github.com/Kjoedicker/blitz/logo"
	"github.com/Kjoedicker/blitz/plan"
	"github.com/Kjoedicker/blitz/request"
	"github.com/Kjoedicker/blitz/results"
)

var (
	testPlan          plan.Plan
	requestPrototypes map[int]request.Request
)

func main() {
	for index := 0; index < len(requestPrototypes); index++ {
		requestPrototype := requestPrototypes[index]

		Execute(requestPrototype)
	}
}

func init() {
	testPlan = plan.Load(cli.TestPlanFilePath)
	requestPrototypes = request.BuildRequestPrototypes(testPlan)

	if cli.PrintLogo {
		logo.Print()
	}
	if cli.PrintResultFormat == "csv" {
		results.PrintCSVHeaders()
	}
}
