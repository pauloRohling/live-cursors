package generator

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNameGenerator_Generate(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != "api-key" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"name": "Paulo"}`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))

	httpClient := http.DefaultClient
	nameGenerator := NewNameGenerator(httpClient, mockServer.URL, "api-key")
	name, err := nameGenerator.Generate()
	assert.NoError(t, err)
	assert.NotEmpty(t, name)
}
