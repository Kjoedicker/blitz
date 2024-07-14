package results

import (
	"log"
	"reflect"

	"github.com/Kjoedicker/blitz/cli"
)

type Field map[string]interface{}
type Fields struct {
	TargetNumber  Field
	RequestGroup  Field
	RequestNumber Field
	ResponseTime  Field
	ErrorResponse Field
}

var fields = Fields{
	TargetNumber: Field{
		"printable": true,
		"header":    "target_number",
	},
	RequestNumber: Field{
		"printable": cli.PrintRequestNumber,
		"header":    "request_number",
	},
	RequestGroup: Field{
		"printable": true,
		"header":    "request_group",
	},
	ResponseTime: Field{
		"printable": true,
		"header":    "response_time",
	},
	ErrorResponse: Field{
		"printable": true,
		"header":    "error_response",
	},
}

func (fields Fields) IsPrintable(field string) bool {
	value := reflect.ValueOf(fields)
	fieldValue := value.FieldByName(field)

	if !fieldValue.IsValid() || fieldValue.Type() != reflect.TypeOf(Field{}) {
		return false
	}

	fieldMap := fieldValue.Interface().(Field)
	printable, exists := fieldMap["printable"]

	return exists && printable.(bool)
}

func (fields Fields) Header(field string) string {
	value := reflect.ValueOf(fields)

	fieldValue := value.FieldByName(field)
	if !fieldValue.IsValid() || fieldValue.Type() != reflect.TypeOf(Field{}) {
		log.Fatalf("Field %s is not valid or not of type Field\n", field)
	}

	fieldMap := fieldValue.Interface().(Field)
	header, exists := fieldMap["header"]
	if !exists {
		log.Fatalf("Field %s does not contain 'header' key\n", field)
	}

	return header.(string)
}
