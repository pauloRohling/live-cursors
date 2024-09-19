package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"live-cursors/internal"
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

	nameGenerator := internal.NewNameGenerator(env.Api.Url, env.Api.Key)
	colorGenerator := internal.NewColorGenerator()

	fmt.Println(env.Api.Url)
	fmt.Println(env.Api.Key)
	fmt.Println(nameGenerator, colorGenerator)

	//http.HandleFunc("/", wsEndpoint)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
