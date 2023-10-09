package main

import (
	"flag"
	"fmt"
	"os"
)

const DEFAULT_INT = 0

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

func enforceRequired(arguments map[string]*int) {
	for description, argument := range arguments {
		if *argument == DEFAULT_INT {
			fmt.Println("Missing required arguments:")
			fmt.Println("\t" + description)
			os.Exit(1)
		}
	}
}

func parseCommandLineArguments() (totalRequests int, interval int) {
	totalRequestsPtr := flag.Int("totalRequests", DEFAULT_INT, "Total requests to be made")
	intervalPtr := flag.Int("interval", DEFAULT_INT, "Interval of when requests are made in seconds")

	flag.Parse()

	arguments := map[string]*int{
		"Total requests: --totalRequests, -t": totalRequestsPtr,
		"Interval: --interval, -i":            intervalPtr,
	}

	enforceRequired(arguments)

	return *totalRequestsPtr, *intervalPtr
}
