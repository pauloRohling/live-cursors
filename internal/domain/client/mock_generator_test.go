package client

type MockGenerator struct {
	result string
}

func NewMockGenerator(result string) *MockGenerator {
	return &MockGenerator{result: result}
}

func (generator MockGenerator) Generate() (string, error) {
	return generator.result, nil
}
