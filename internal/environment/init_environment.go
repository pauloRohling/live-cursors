package environment

import (
	"github.com/ilyakaznacheev/cleanenv"
	"live-cursors/pkg/banner"
)

var env Config

func Env() Config {
	return env
}

func init() {
	banner.Show()
	if err := cleanenv.ReadConfig("env.yml", &env); err != nil {
		panic(err)
	}
}
