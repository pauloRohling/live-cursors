package generator

type Generator[T any] interface {
	Generate() (T, error)
}
