package domain

import "live-cursors/internal/model"

type UserGenerator struct {
	nameGenerator  Generator[string]
	colorGenerator Generator[string]
}

func NewUserGenerator(nameGenerator Generator[string], colorGenerator Generator[string]) *UserGenerator {
	return &UserGenerator{nameGenerator: nameGenerator, colorGenerator: colorGenerator}
}

func (generator *UserGenerator) Generate() (*model.User, error) {
	name, err := generator.nameGenerator.Generate()
	if err != nil {
		return nil, err
	}

	color, err := generator.colorGenerator.Generate()
	if err != nil {
		return nil, err
	}

	return model.NewUser(name, color), nil
}
