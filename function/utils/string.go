package utils

import (
	"strconv"
)

func StringToFloat64(str string) float64 {

	var result float64

	result, _ = strconv.ParseFloat(str, 64)

	return result
}

func StringSplit(str string, length int) string {

	var result string

	if len(str)+1 >= length {
		result = str[:length-1] + ".."
	} else {
		result = str
		for i := 0; i < length-(len(str)+1); i++ {
			result = result + " "
		}
	}

	return result
}

func BoolStringToFloat64(str string) float64 {

	var result float64

	if str == "true" {
		result = 1
	}

	return result
}
