package main

import (
	"github.com/Kjoedicker/blitz/cli"
	_ "github.com/Kjoedicker/blitz/logo"
	"github.com/Kjoedicker/blitz/plan"
	"github.com/Kjoedicker/blitz/request"
)

func main() {
	testPlan := plan.Load(cli.TestPlanFilePath)
	requestPrototypes := request.BuildRequestPrototypes(testPlan)

	for index := 0; index < len(requestPrototypes); index++ {
		requestPrototype := requestPrototypes[index]

		Execute(requestPrototype)
	}
}
