package insert

import (
	"fmt"
	"strconv"
)

type Call struct {
	BaseUrl string
	Route   string
	Params  map[string]string
}

type Resource struct {
	Call
	Name  string
	Count int
}

func (r *Resource) SetCount() error {
	client := &Client{}
	client.SetBaseUrl(r.BaseUrl)

	res, err := client.Get(r.Route, r.Params, Params{
		"@count": "true",
	})

	if err != nil {
		r.Count = 0
		return fmt.Errorf("error setting count: %w", err)
	}

	count, err := strconv.Atoi(res)

	if err != nil {
		r.Count = 0
		return fmt.Errorf("error converting count to int: %w", err)
	}

	r.Count = count
	return nil
}
