package internal

import "math/rand"

type NameGenerator struct {
	url string
	key string
}

func NewNameGenerator(url string, key string) *NameGenerator {
	return &NameGenerator{
		url: url,
		key: key,
	}
}

func (n *NameGenerator) Generate() string {
	names := []string{"Paulo", "João", "Pedro", "Maria", "José", "Carlos", "Ana", "Maria", "Luís"}
	index := rand.Intn(len(names))
	return names[index]
}
