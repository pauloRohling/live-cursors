package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"live-cursors/internal/domain"
)

type Environment struct {
	Api struct {
		Url string `yml:"url" env:"API_URL"`
		Key string `yml:"key" env:"API_KEY"`
	} `yml:"api"`
}

var env Environment

func main() {
	err := cleanenv.ReadConfig("env.yml", &env)
	if err != nil {
		panic(err)
	}

	nameGenerator := domain.NewNameGenerator(env.Api.Url, env.Api.Key)
	colorGenerator := domain.NewColorGenerator()
	userGenerator := domain.NewUserGenerator(nameGenerator, colorGenerator)

	user, err := userGenerator.Generate()
	if err != nil {
		panic(err)
	}

	fmt.Println(env.Api.Url)
	fmt.Println(env.Api.Key)
	fmt.Println(user)

	//http.HandleFunc("/", wsEndpoint)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
