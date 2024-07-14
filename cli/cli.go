package cli

import (
	"flag"
	"log"
)

const (
	DEFAULT_FILE_PATH                   = "./test-plan.yml"
	DEFAULT_PRINT_RESULT_FORMAT         = "csv"
	DEFAULT_ENABLE_PRINT_LOGO           = true
	DEFAULT_ENABLE_PRINT_REQUEST_NUMBER = false
)

var (
	TestPlanFilePath   string
	PrintResultFormat  string
	PrintLogo          bool
	PrintRequestNumber bool
)

func parseFlags() {
	// Define flags
	flag.StringVar(
		&TestPlanFilePath,
		"test-plan-path",
		DEFAULT_FILE_PATH,
		"Test plan path. Defaults to: "+DEFAULT_FILE_PATH,
	)
	flag.StringVar(
		&PrintResultFormat,
		"format",
		DEFAULT_PRINT_RESULT_FORMAT,
		"The format the results are printed in. Defaults to: "+DEFAULT_PRINT_RESULT_FORMAT,
	)
	flag.BoolVar(
		&PrintLogo,
		"print-logo",
		DEFAULT_ENABLE_PRINT_LOGO,
		"Enable/Disable the printed logo. Defaults to: true",
	)
	flag.BoolVar(
		&PrintRequestNumber,
		"print-request-number",
		DEFAULT_ENABLE_PRINT_REQUEST_NUMBER,
		"Enable/Disable printing each request. The stops calculating the mean of all the requests in a group. Defaults to: false",
	)

	// Parse flags
	flag.Parse()
}

func validatePrintResultFormat() {
	allowedFormats := map[string]bool{
		"csv":  true,
		"text": true,
	}
	if !allowedFormats[PrintResultFormat] {
		log.Fatalf("Invalid format: %s. Allowed values are: csv, text.\n", PrintResultFormat)
	}
}

func validateFlags() {
	validatePrintResultFormat()
}

func init() {
	parseFlags()
	validateFlags()
}
