package results

import (
	"fmt"
	"strings"

	"github.com/Kjoedicker/blitz/cli"
	"github.com/Kjoedicker/blitz/request"
)

func PrintAll(requests request.Requests) {
	for _, request := range requests {
		switch cli.PrintResultFormat {
		case "text":
			Print(request)
		default:
			PrintToCSV(request)
		}
	}
}

func Print(request request.Request) {
	fmt.Printf(
		"Test Plan Number: %d "+
			"Request Group: %d "+
			"Request Number %d "+
			"Response Time: %f "+
			"Errors: %v \n",
		request.TestPlanNumber,
		request.RequestGroup,
		request.RequestNumber,
		request.ResponseTime,
		request.ErrorResponse,
	)
}

func PrintToCSV(request request.Request) {
	fmt.Printf(
		"%d,%d,%d,%f,%v\n",
		request.TestPlanNumber,
		request.RequestGroup,
		request.RequestNumber,
		request.ResponseTime,
		request.ErrorResponse,
	)
}

func PrintCSVHeaders() {
	csvHeaders := strings.Join(
		[]string{
			"test_plan_number",
			"request_group",
			"request_number",
			"response_time",
			"error_response",
		},
		",",
	)

	fmt.Println(csvHeaders)
}
