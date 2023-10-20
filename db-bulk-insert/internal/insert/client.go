package insert

import (
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Params map[string]string

type Client struct {
	baseUrl string
}

func (c *Client) BaseUrl() string {
	return c.baseUrl
}

func (c *Client) SetBaseUrl(url string) {
	url, _ = strings.CutSuffix(url, "/")

	c.baseUrl = url
}

func (c *Client) Get(route string, params ...Params) (string, error) {
	err := c.validateRoute(route)

	if err != nil {
		return "", err
	}

	err = c.validateParams(params...)

	if err != nil {
		return "", err
	}

	formattedParams := c.formatParams(params...)

	requestURL := c.baseUrl + route + formattedParams

	response, err := http.Get(requestURL)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyBuffer, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(bodyBuffer), nil
}

func (c *Client) validateRoute(route string) error {

	if route == "" {
		return nil
	}

	validPattern := regexp.MustCompile(`^(\/[a-zA-Z0-9]+)*\/[a-zA-Z0-9]+(\/|)$`)

	if !validPattern.MatchString(route) {
		return errors.New("invalid route provided")
	}

	return nil
}

func (c *Client) validateParams(params ...Params) error {
	for _, param := range params {
		for key, value := range param {
			if key == "" || value == "" {
				return errors.New("invalid param provided")
			}
		}
	}

	return nil
}

func (c *Client) formatParams(params ...Params) string {
	formattedParams := ""

	for _, param := range params {
		for key, value := range param {
			formattedParams += "&" + key + "=" + value
		}
	}

	if formattedParams != "" {
		formattedParams = formattedParams[1:]
		formattedParams = "?" + formattedParams
	}

	return formattedParams
}
