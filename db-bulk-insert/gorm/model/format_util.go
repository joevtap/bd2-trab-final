package model

import "fmt"

func Format(array []int32) string {
	if len(array) == 0 {
		return "{}"
	}

	str := "{"

	for _, element := range array {
		str += fmt.Sprintf("%d,", element)
	}

	str = str[:len(str)-1]
	str += "}"

	return str
}
