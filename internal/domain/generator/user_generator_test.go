package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserGenerator_Generate(t *testing.T) {
	nameGenerator := NewNameGenerator("https://api.livecursors.com", "api-key")
	colorGenerator := NewColorGenerator()
	userGenerator := NewUserGenerator(nameGenerator, colorGenerator)

	user, err := userGenerator.Generate()
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Color)
}
