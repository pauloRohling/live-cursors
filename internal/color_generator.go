package internal

import (
	"fmt"
	"math/rand/v2"
)

type ColorGenerator struct {
}

func NewColorGenerator() *ColorGenerator {
	return &ColorGenerator{}
}

func (c *ColorGenerator) Generate() string {
	return fmt.Sprintf("#%x%x%x", rand.UintN(255), rand.UintN(255), rand.UintN(255))
}
