package plan

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Description string              `yaml:"description"`
	Method      string              `yaml:"method"`
	Path        string              `yaml:"path"`
	Headers     map[string][]string `yaml:"headers"`
	Hits        int                 `yaml:"hits"`
	Interval    string              `yaml:"interval"`
	Duration    int                 `yaml:"duration"`
}

type Plan struct {
	Host    string   `yaml:"host"`
	Targets []Target `yaml:"Targets"`
}

func (target Target) PrintDetails() {
	fmt.Println("\nScenario: " + target.Description)
	fmt.Printf("%d requests every %s for %d minutes\n", target.Hits, target.Interval, target.Duration)
}

func Load(filePath string) Plan {
	plan := &Plan{}

	planFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("os.ReadFile %v ", err)
	}

	err = yaml.Unmarshal(planFile, plan)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return *plan
}
