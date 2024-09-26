package domain

type Generator[T any] interface {
	Generate() (T, error)
}
