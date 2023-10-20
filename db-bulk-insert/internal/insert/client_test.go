package insert

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	getMockApi := func() *httptest.Server {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseBody := `{"message":"ok"}`

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", fmt.Sprint(len(responseBody)))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(responseBody))

		}))

		return mockServer
	}

	mockApi := getMockApi()
	defer mockApi.Close()

	t.Run("should validate route", func(t *testing.T) {
		client := &Client{baseUrl: mockApi.URL}

		testCases := []struct {
			route string
		}{
			{" "},
			{"invalid"},
			{"invalid/"},
			{"invalid/invalid"},
			{"invalid/invalid/"},
			{"/invalid//invalid/invalid"},
		}

		for _, testCase := range testCases {
			err := client.validateRoute(testCase.route)

			assert.Error(t, err)
		}

		err := client.validateRoute("")
		assert.NoError(t, err)
	})

	t.Run("should error if invalid route is provided", func(t *testing.T) {
		client := &Client{baseUrl: mockApi.URL}
		testCases := []struct {
			route string
		}{
			{route: " "},
			{route: "invalid"},
			{route: "invalid/"},
		}

		for _, testCase := range testCases {
			_, err := client.Get(testCase.route, Params{})

			assert.Error(t, err)
		}
	})

	t.Run("should return response body", func(t *testing.T) {
		client := &Client{baseUrl: mockApi.URL}

		actual, err := client.Get("", Params{})

		expected := `{"message":"ok"}`

		t.Log(actual)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should be able to make multiple requests", func(t *testing.T) {
		client := &Client{baseUrl: mockApi.URL}

		for i := 0; i < 10; i++ {
			expected := `{"message":"ok"}`
			actual, err := client.Get("", Params{})
			assert.NoError(t, err)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("should be able to take in params", func(t *testing.T) {
		client := &Client{baseUrl: mockApi.URL}

		actual, err := client.Get("", Params{"param1": "true", "param2": "value2"})
		expected := `{"message":"ok"}`

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should format params correctly", func(t *testing.T) {
		client := &Client{baseUrl: mockApi.URL}

		actual := client.formatParams(Params{"test": "true", "test2": "somevalue"})
		expected1 := "?test=true&test2=somevalue"
		expected2 := "?test2=somevalue&test=true"

		assert.True(t, actual == expected1 || actual == expected2)
	})

	t.Run("should error if api call fails", func(t *testing.T) {
		client := &Client{baseUrl: "http://invalid.url"}

		_, err := client.Get("", Params{})

		assert.Error(t, err)
	})
}
