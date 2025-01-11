package customValidators

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
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

func ParseGender(genderStr string) (pgtype.Text, bool) {
	if genderStr == "" {
		return pgtype.Text{Valid: false, String: ""}, true
	}

	_, isValid := validGender[genderStr]
	return pgtype.Text{String: genderStr, Valid: true}, isValid
}

func ParseDepartmentID(id string) (*int, bool) {
	if id == "" {
		return nil, true
	}
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		return nil, false
	}

	if departmentID < 1 {
		return nil, false
	}

	return &departmentID, true
}
