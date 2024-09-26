package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorGenerator_Generate(t *testing.T) {
	colorGenerator := NewColorGenerator()
	color, err := colorGenerator.Generate()
	assert.NoError(t, err)
	assert.NotEmpty(t, color)
	assert.Len(t, color, 7)
	assert.EqualValues(t, color[0], '#')
}
