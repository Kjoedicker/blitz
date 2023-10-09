package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type TestPlan struct {
	Host    string `yaml:"host"`
	Targets []struct {
		Description string              `yaml:"description"`
		Method      string              `yaml:"method"`
		Path        string              `yaml:"path"`
		Headers     map[string][]string `yaml:"headers"`
		Hits        int                 `yaml:"hits"`
		Interval    int                 `yaml:"interval"`
	} `yaml:"Targets"`
}

func loadTestPlan(filePath string) TestPlan {
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
