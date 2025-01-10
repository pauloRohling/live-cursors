package environment

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

var env Config

func Env() Config {
	return env
}

func init() {
	showBanner()
	if err := cleanenv.ReadConfig("env.yml", &env); err != nil {
		panic(err)
	}
}

func showBanner() {
	if file, err := os.ReadFile("./banner.txt"); err == nil {
		fmt.Println(string(file))
	}
}
