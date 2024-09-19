package internal

type UserGenerator struct {
}

func NewUserGenerator() *UserGenerator {
	return &UserGenerator{}
}

func (u *UserGenerator) Generate() string {
	return "user"
}
