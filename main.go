package main

import (
	"github.com/Kjoedicker/blitz/cli"
	"github.com/Kjoedicker/blitz/logo"
	"github.com/Kjoedicker/blitz/plan"
	"github.com/Kjoedicker/blitz/request"
)

func main() {
	logo.Print()
	planFilePath := cli.ParseTestPlanPath()
	testPlan := plan.Load(planFilePath)
	requests := request.BuildRequests(testPlan)

	for _, requestPrototype := range requests {
		Execute(requestPrototype)
	}
}
