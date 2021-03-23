package utils

import (
	"strconv"
)

func F64ToString(float float64) string {
	return strconv.FormatFloat(float, 'f', 4, 64)
}
