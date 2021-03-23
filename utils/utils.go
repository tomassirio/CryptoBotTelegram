package utils

import (
	"strconv"
)

func F64ToString(float float64) string {
	return strconv.FormatFloat(float, 'f', 2, 64)
}

func BoldString(string string) string {
	return "**" + string + "**"
}