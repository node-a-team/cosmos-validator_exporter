package utils

import ()

func BoolToFloat64(b bool) float64 {

	var result float64

	if b == true {
		result = 1
	} else {
		result = 0
	}

	return result
}
