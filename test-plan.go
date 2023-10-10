package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Description string              `yaml:"description"`
	Method      string              `yaml:"method"`
	Path        string              `yaml:"path"`
	Headers     map[string][]string `yaml:"headers"`
	Hits        int                 `yaml:"hits"`
	Interval    int                 `yaml:"interval"`
	Duration    int                 `yaml:"duration"`
}

type TestPlan struct {
	Host    string   `yaml:"host"`
	Targets []Target `yaml:"Targets"`
}

func (target Target) PrintDetails() {
	fmt.Println("\nScenario: " + target.Description)
	fmt.Printf("%d requests every %d minutes for %d minutes\n", target.Hits, target.Interval, target.Duration)
}

func LoadTestPlan(filePath string) TestPlan {
	testPlan := &TestPlan{}

	testPlanFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("os.ReadFile %v ", err)
	}

	err = yaml.Unmarshal(testPlanFile, testPlan)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return *testPlan
}

func (testPlan TestPlan) Begin() {
	for _, target := range testPlan.Targets {
		target.PrintDetails()

		rawUrl, err := url.JoinPath(testPlan.Host, target.Path)
		if err != nil {
			log.Fatal(err)
		}

		request := Request{
			Method:  target.Method,
			Headers: target.Headers,
			rawUrl:  rawUrl,
		}.Init()

		hitsPerSecond := HitsPer(target.Hits, MinutesToSeconds(target.Interval))
		MakeRequestsForDuration(request, target.Duration, hitsPerSecond)
	}
}
