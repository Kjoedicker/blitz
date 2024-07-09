package plan

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

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

func validateRequiredField[T any](v T) bool {
	var zero T // Create a default value of type T to check against
	return !reflect.DeepEqual(v, zero)
}

func (plan *Plan) validate() {
	for _, target := range plan.Targets {
		fields := map[string]interface{}{
			"Description": target.Description,
			"Path":        target.Path,
			"Method":      target.Method,
			"Hits":        target.Hits,
			"Interval":    target.Interval,
			"Duration":    target.Duration,
		}

		var validationFailures []string

		for fieldName, data := range fields {
			valid := validateRequiredField(data)
			if !valid {
				validationFailures = append(validationFailures, fmt.Sprintf("Key: \"%s\" Value: %v \n", fieldName, data))
			}
		}

		if len(validationFailures) > 0 {
			log.Fatalf("Plan validation failed:\n\t%v", strings.Join(validationFailures, "\t"))
		}
	}
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

	plan.validate()

	return *plan
}
