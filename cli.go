package main

import (
	"flag"
	"fmt"
)

const (
	DEFAULT_FILE_PATH = "./test-plan.yml"
)

func printLogo() {
	blitzLogo := `
▄▄▄▄   ██▓    ██▄▄▄█████▒███████▒
▓█████▄▓██▒   ▓██▓  ██▒ ▓▒ ▒ ▒ ▄▀░
▒██▒ ▄█▒██░   ▒██▒ ▓██░ ▒░ ▒ ▄▀▒░ 
▒██░█▀ ▒██░   ░██░ ▓██▓ ░  ▄▀▒   ░
░▓█  ▀█░██████░██░ ▒██▒ ░▒███████▒
░▒▓███▀░ ▒░▓  ░▓   ▒ ░░  ░▒▒ ▓░▒░▒
▒░▒   ░░ ░ ▒  ░▒ ░   ░   ░░▒ ▒ ░ ▒
 ░    ░  ░ ░   ▒ ░ ░     ░ ░ ░ ░ ░
 ░         ░  ░░           ░ ░    
	  ░                  ░        
`

	fmt.Println(blitzLogo)
}

func parseTestPlanPath() (testPlanFilePath string) {
	testPlanFilePathPtr := flag.String("test-plan-path", DEFAULT_FILE_PATH, "Test plan path. Defaults to: "+DEFAULT_FILE_PATH)

	flag.Parse()

	return *testPlanFilePathPtr
}
