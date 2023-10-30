package model

type TermsType struct {
	Tag       []string `json:"tag"`
	Linguagem []string `json:"linguagem"`
	Area      []string `json:"area"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timestamp struct {
	Date         string `json:"date"`
	TimezoneType int32  `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type Model interface {
	TableName() string
	Formatter
}
