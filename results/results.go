package results

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Kjoedicker/blitz/cli"
	"github.com/Kjoedicker/blitz/request"
)

const (
	FormatText = "text"
	FormatCSV  = "csv"
)

func formatColumn(field string, value interface{}) string {
	if cli.PrintResultFormat == FormatText {
		return fmt.Sprintf("%s: %v", field, value)
	}
	return fmt.Sprintf("%v", value)
}

func getColumnDelimiter() string {
	if cli.PrintResultFormat == FormatText {
		return " "
	}
	return ","
}

func formatRequestColumns(request interface{}) string {
	var (
		reqValues = reflect.ValueOf(request)
		reqType   = reqValues.Type()

		columnDelimiter = getColumnDelimiter()
		columns         = []string{}
	)
	for i := 0; i < reqValues.NumField(); i++ {
		field := reqType.Field(i).Name
		value := reqValues.Field(i)

		if fields.IsPrintable(field) {
			columns = append(columns, formatColumn(field, value))
		}
	}
	return strings.Join(columns, columnDelimiter)
}

func Print(request request.Request) {
	rows := formatRequestColumns(request)
	fmt.Println(rows)
}

func PrintAll(requests request.Requests) {
	for _, request := range requests {
		Print(request)
	}
}

func PrintCSVHeaders() {
	var (
		headers     = []string{}
		fieldValues = reflect.ValueOf(fields)
		fieldType   = fieldValues.Type()
	)
	for i := 0; i < fieldValues.NumField(); i++ {
		field := fieldType.Field(i).Name

		if fields.IsPrintable(field) {
			headers = append(headers, fields.Header(field))
		}
	}
	fmt.Println(strings.Join(headers, ","))
}
