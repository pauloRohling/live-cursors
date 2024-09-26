package client

type Generator[T any] interface {
	Generate() (T, error)
}
