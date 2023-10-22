package insert

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Call struct {
	BaseUrl  string
	Route    string
	Params   Params
	Limit    int
	CurrPage int
}

type Resource struct {
	Call
	Name  string
	Count int
	Model interface{}
}

func (r *Resource) SetCount() error {
	client := &Client{}
	client.SetBaseUrl(r.BaseUrl)

	res, err := client.Get(r.Route, Params{
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

func (r *Resource) SetPage(page int) {
	r.CurrPage = page
}

func (r *Resource) SetLimit(limit int) {
	r.Limit = limit
}

func (r *Resource) GetMany() (string, error) {
	client := &Client{}
	client.SetBaseUrl(r.BaseUrl)

	res, err := client.Get(r.Route, r.Params, Params{
		"@limit": strconv.Itoa(r.Limit),
		"@page":  strconv.Itoa(r.CurrPage),
	})

	if err != nil {
		return "", fmt.Errorf("error getting resource: %w", err)
	}

	return res, nil
}

func GetBatch[T any](resource Resource, sleep time.Duration, ch chan []T) error {
	client := &Client{}
	client.SetBaseUrl(resource.BaseUrl)

	resource.SetCount()

	var results []T

	waitChan := make(chan struct{})
	errorChan := make(chan error)

	for i := 1; i <= resource.Count; i += resource.Limit {
		go func(i int, results *[]T, sleep time.Duration) {
			log.Printf("Getting records from page %d/%d of resource %q\n", i/resource.Limit+1, (resource.Count+resource.Limit-1)/resource.Limit, resource.Name)

			res, err := client.Get(resource.Route, resource.Params, Params{
				"@limit": strconv.Itoa(resource.Limit),
				"@page":  strconv.Itoa(i/resource.Limit + 1),
			})

			if err != nil {
				errorChan <- fmt.Errorf("error getting resource: %w", err)
				return
			}

			var responses []T
			if err := json.Unmarshal([]byte(res), &responses); err != nil {
				errorChan <- fmt.Errorf("error unmarshaling response: %w", err)
				return
			}

			errorChan <- nil

			*results = append(*results, responses...)

			ch <- responses
			waitChan <- struct{}{}
		}(i, &results, sleep)

		if err := <-errorChan; err != nil {
			return err
		}

		time.Sleep(sleep * time.Millisecond)
	}

	for i := 1; i <= resource.Count; i += resource.Limit {
		<-waitChan
	}

	return nil
}

func UnmarshalResponse[T any](res string, err error) (T, error) {
	var ret T

	if err != nil {
		return ret, fmt.Errorf("error getting resource: %w", err)
	}

	err = json.Unmarshal([]byte(res), &ret)

	if err != nil {
		return ret, fmt.Errorf("error unmarshaling resource response: %w", err)
	}

	return ret, nil
}
