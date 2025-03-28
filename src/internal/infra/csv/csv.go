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

// func Convert(data interface{}) string {
// 	var valueConverted string = ""
// 	switch v := data.(type) {
// 	case int64, int32, int16, int8, int:
// 		valueConverted = fmt.Sprintf("%d", v)
// 	// case float64, float32:
// 	// 	valueConverted = fmt.Sprintf("%f", v)
// 	case bool:
// 		valueConverted = fmt.Sprintf("%t", v)
// 	case []byte:
// 		valueConverted = string(v)
// 	case string:
// 		valueConverted = v
// 	case time.Time:
// 		valueConverted = v.Format(time.RFC3339)
// 	default:
// 		valueConverted = fmt.Sprintf("%v", v)
// 	}
// 	if data == nil {
// 		return ""
// 	}
// 	return valueConverted
// }
