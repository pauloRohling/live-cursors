package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/domain/generator"
	"live-cursors/internal/presentation"
	"log"
	"net/http"
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

	nameGenerator := generator.NewNameGenerator(env.Api.Url, env.Api.Key)
	colorGenerator := generator.NewColorGenerator()
	clientManager := client.NewInMemoryManager()
	clientFactory := client.NewRandomFactory(nameGenerator, colorGenerator)

	wsHandler := presentation.NewWebSocketHandler(clientFactory, clientManager)

	http.HandleFunc("/", wsHandler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
