package strutil

import (
	"fmt"
	"regexp"
	"time"
)

const DATE_FORMAT = "2006-01-02 15:04:05"

func RemoveSpecialCodes(valueData string) string {
	re := regexp.MustCompile(`[\n\t\r]`)
	valueData = re.ReplaceAllString(valueData, " ")
	return valueData
}

// func RemoveCSVUnsafeChars(valueData string) string {
// 	re := regexp.MustCompile(`[^\w\s,.-]`)
// 	valueData = re.ReplaceAllString(valueData, " ")
// 	return valueData
// }

func ConvertToDateTime(valueData string) (string, error) {
	formats := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999",
		"2006-01-02 15:04:05.999999",
		"02/01/2006",
		"02/01/2006 15:04:05",
		"02/01/2006 15:04:05.999",
		"02/01/2006 15:04:05.999999",
		"02-01-2006",
		"02-01-2006 15:04:05",
		"02-01-2006 15:04:05.999",
		"02-01-2006 15:04:05.999999",
	}
	for _, format := range formats {
		if t, err := time.Parse(format, valueData); err == nil {
			return t.Format(DATE_FORMAT), nil
		}
	}
	return valueData, fmt.Errorf("☠️ Error: Date %s not recognized", valueData)
}

func ConvertFloatToString(valueData string, decimalSeparator string) (string, error) {
	decimalPointIndex := regexp.MustCompile(`\.`).FindStringIndex(valueData)
	if decimalPointIndex == nil {
		return valueData, nil
	}

	var intValue string = valueData[:decimalPointIndex[0]]
	var decimalValue string = valueData[decimalPointIndex[0]+1:]

	retorno := fmt.Sprintf("%s%s%s", intValue, decimalSeparator, decimalValue)
	return retorno, nil
}
