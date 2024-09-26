package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameGenerator_Generate(t *testing.T) {
	nameGenerator := NewNameGenerator("https://api.livecursors.com", "api-key")
	name, err := nameGenerator.Generate()
	assert.NoError(t, err)
	assert.NotEmpty(t, name)
}
