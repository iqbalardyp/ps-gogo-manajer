package customValidators

import (
	"strconv"
)

var validGender = map[string]bool{
	"male":   true,
	"female": true,
}

func ParseLimitOffset(val string, defaultVal int) int {
	parsedVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}

	if parsedVal < 0 {
		return defaultVal
	}

	return parsedVal
}

func ParseGender(genderStr string) (string, bool) {
	if genderStr == "" {
		return "", true
	}

	_, isValid := validGender[genderStr]
	return genderStr, isValid
}

func ParseDepartmentID(id string) (int, bool) {
	if id == "" {
		return 0, true
	}
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		return 0, false
	}

	if departmentID < 1 {
		return 0, false
	}

	return departmentID, true
}
