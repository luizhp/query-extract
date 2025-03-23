package strutil

import (
	"fmt"
	"regexp"
	"strconv"
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
	floatValueConverted, err := strconv.ParseFloat(valueData, 64)
	if err != nil {
		return valueData, err
	}

	stringValueConverted := strconv.FormatFloat(floatValueConverted, 'f', -1, 64)
	re := regexp.MustCompile(`\.(\d*)`)
	matches := re.FindStringSubmatch(stringValueConverted)
	var decimalPart string = ""
	if len(matches) > 1 {
		decimalPart = matches[1]
	}

	intPart := int64(floatValueConverted)
	if decimalPart == "" {
		decimalSeparator = ""
	}

	return fmt.Sprintf("%d%s%s", intPart, decimalSeparator, decimalPart), nil
}
