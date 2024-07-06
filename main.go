package main

import (
	"fmt"

	"github.com/Kjoedicker/blitz/cli"
	"github.com/Kjoedicker/blitz/logo"
	"github.com/Kjoedicker/blitz/plan"
	"github.com/Kjoedicker/blitz/request"
)

func main() {
	logo.Print()

	planFilePath := cli.ParseTestPlanPath()
	testPlan := plan.Load(planFilePath)
	requestPrototypes := request.BuildRequestPrototypes(testPlan)

	for index := 0; index < len(requestPrototypes); index++ {
		requestPrototype := requestPrototypes[index]

		fmt.Printf("Executing test plan: %d", index+1)

		Execute(requestPrototype)
	}
}
