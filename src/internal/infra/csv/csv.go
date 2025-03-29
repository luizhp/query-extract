package csv

import (
	"bytes"
)

const COLUMN_DELIMITER = ";" // "\t"
const NEW_LINE = "\n"

const DECIMAL_SEPARATOR = ","

func Generate(columns []string, rows []map[string]string) string {
	return header(columns) + detail(columns, rows)
}

func header(columns []string) string {
	var buffer bytes.Buffer
	for _, colName := range columns {
		buffer.WriteString(colName)
		buffer.WriteString(COLUMN_DELIMITER)
	}
	buffer.WriteString(NEW_LINE)
	return buffer.String()
}

func detail(columns []string, rows []map[string]string) string {
	var buffer bytes.Buffer
	for _, row := range rows {
		for _, colName := range columns {
			buffer.WriteString(row[colName])
			buffer.WriteString(COLUMN_DELIMITER)
		}
		buffer.WriteString(NEW_LINE)
	}
	return buffer.String()
}
