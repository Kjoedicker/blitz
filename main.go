package main

import (
	"github.com/Kjoedicker/blitz/cli"
	"github.com/Kjoedicker/blitz/logo"
	"github.com/Kjoedicker/blitz/plan"
	"github.com/Kjoedicker/blitz/request"
	"github.com/Kjoedicker/blitz/results"
)

var (
	testPlan plan.Plan
	targets  request.Requests
)

func main() {
	Setup(targets)
}

func init() {
	testPlan = plan.Load(cli.TestPlanFilePath)
	targets = request.BuildTargets(testPlan)

	if cli.PrintLogo {
		logo.Print()
	}
	if cli.PrintResultFormat == "csv" {
		results.PrintCSVHeaders()
	}
}
