package main

import "fmt"

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

func main() {
	printLogo()

	testPlanPath := ParseTestPlanPath()
	testPlan := LoadTestPlan(testPlanPath)

	testPlan.Begin()
}
