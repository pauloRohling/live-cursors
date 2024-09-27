package main

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/ilyakaznacheev/cleanenv"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/domain/generator"
	"live-cursors/internal/domain/message"
	"live-cursors/internal/presentation"
	"log"
	"net/http"
	"time"
)

type Environment struct {
	Api struct {
		Url string `yml:"url" env:"API_URL"`
		Key string `yml:"key" env:"API_KEY"`
	} `yml:"api"`
	Http struct {
		MaxRetry        int           `yml:"max_retry" env:"HTTP_MAX_RETRY"`
		MaxRetryTimeout time.Duration `yml:"max_retry_timeout" env:"HTTP_MAX_RETRY_TIMEOUT"`
		MinRetryTimeout time.Duration `yml:"min_retry_timeout" env:"HTTP_MIN_RETRY_TIMEOUT"`
	} `yml:"http"`
}

var env Environment

func main() {
	err := cleanenv.ReadConfig("env.yml", &env)
	if err != nil {
		panic(err)
	}

	httpClient := GetHttpClient()
	nameGenerator := generator.NewNameGenerator(httpClient, env.Api.Url, env.Api.Key)
	colorGenerator := generator.NewColorGenerator()
	clientManager := client.NewInMemoryManager()
	clientFactory := client.NewRandomFactory(nameGenerator, colorGenerator)
	producer := message.NewProducer(clientManager)

	wsHandler := presentation.NewWebSocketHandler(clientFactory, clientManager, producer)

	http.HandleFunc("/", wsHandler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetHttpClient() *http.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = env.Http.MaxRetry
	httpClient.RetryWaitMax = env.Http.MaxRetryTimeout
	httpClient.RetryWaitMin = env.Http.MinRetryTimeout
	httpClient.Logger = log.Default()
	return httpClient.StandardClient()
}
