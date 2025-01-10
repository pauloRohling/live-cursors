package generator

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NameGenerator struct {
	httpClient *http.Client
	url        string
	key        string
}

func NewNameGenerator(httpClient *http.Client, url string, key string) *NameGenerator {
	return &NameGenerator{
		httpClient: httpClient,
		url:        url,
		key:        key,
	}
}

func (generator *NameGenerator) Generate() (string, error) {
	request, err := http.NewRequest("GET", generator.url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("X-API-KEY", generator.key)

	response, err := generator.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var payload struct {
		Name string `json:"name"`
	}

	if err = json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return "", err
	}

	return payload.Name, nil
}
