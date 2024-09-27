package generator

import (
	"fmt"
	"math/rand/v2"
)

type ColorGenerator struct {
}

func NewColorGenerator() *ColorGenerator {
	return &ColorGenerator{}
}

func (generator *ColorGenerator) Generate() (string, error) {
	return fmt.Sprintf(
		"#%02X%02X%02X",
		rand.UintN(255),
		rand.UintN(255),
		rand.UintN(255),
	), nil
}
