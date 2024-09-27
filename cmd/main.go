package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/ilyakaznacheev/cleanenv"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/domain/generator"
	"live-cursors/internal/domain/message"
	"live-cursors/internal/presentation"
	"live-cursors/pkg/graceful"
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
	server := &http.Server{Addr: fmt.Sprintf(":%d", 8080)}

	gracefulShutdownCtx := graceful.Shutdown(&graceful.Params{
		OnStart:   func() { log.Printf("Graceful shutdown started. Waiting for active requests to complete") },
		OnTimeout: func() { log.Fatal("Graceful shutdown timed out. Forcing exit.") },
		OnShutdown: func(timeoutCtx context.Context) {
			if shutdownErr := server.Shutdown(timeoutCtx); shutdownErr != nil {
				log.Fatal(shutdownErr.Error())
			}
		},
	})

	log.Printf("Web server started listening on post %d", 8080)
	if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Could not initialize web server on port %d", 8080)
	}

	<-gracefulShutdownCtx.Done()
	log.Println("Graceful shutdown complete")
}

func GetHttpClient() *http.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = env.Http.MaxRetry
	httpClient.RetryWaitMax = env.Http.MaxRetryTimeout
	httpClient.RetryWaitMin = env.Http.MinRetryTimeout
	httpClient.Logger = log.Default()
	return httpClient.StandardClient()
}
