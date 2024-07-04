package cli

import (
	"flag"
)

const (
	DEFAULT_FILE_PATH = "./test-plan.yml"
)

func ParseTestPlanPath() (testPlanFilePath string) {
	testPlanFilePathPtr := flag.String("target-path", DEFAULT_FILE_PATH, "Test plan path. Defaults to: "+DEFAULT_FILE_PATH)

	flag.Parse()

	return *testPlanFilePathPtr
}
