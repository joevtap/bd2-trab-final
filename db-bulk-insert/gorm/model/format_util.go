package model

import (
	"fmt"
)

type Formatter interface {
	Format()
}

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

func FormatTimestamp(timestamp Timestamp) string {
	return timestamp.Date
}

func FormatQuery(t TermsType) string {
	if len(t.Tag) == 0 && len(t.Linguagem) == 0 {
		return ""
	}

	str := ""

	for _, element := range t.Tag {
		str += fmt.Sprintf("%s ", element)
	}

	for _, element := range t.Linguagem {
		str += fmt.Sprintf("%s ", element)
	}

	str = str[:len(str)-1]

	return str
}
