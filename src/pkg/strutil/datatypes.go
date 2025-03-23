package strutil

import "strconv"

func IsInteger(value string) bool {
	if _, err := strconv.Atoi(value); err != nil {
		return false
	}
	return true
}

func IsFloat(value string) bool {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return false
	}
	return true
}

func IsBoolean(value string) bool {
	if _, err := strconv.ParseBool(value); err != nil {
		return false
	}
	return true
}

func IsDate(value string) bool {
	if _, err := ConvertToDateTime(value); err != nil {
		return false
	}
	return true
}
