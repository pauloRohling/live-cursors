package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	clientApplication "live-cursors/internal/application/client"
	"live-cursors/internal/application/generator"
	"live-cursors/internal/application/message"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/environment"
	"live-cursors/internal/presentation"
	"live-cursors/pkg/graceful"
	"log"
	"net/http"
)

var env = environment.Env()

func main() {
	httpClient := GetHttpClient()
	nameGenerator := generator.NewNameGenerator(httpClient, env.Api.Url, env.Api.Key)
	colorGenerator := generator.NewColorGenerator()
	clientManager := clientApplication.NewInMemoryManager()
	clientFactory := client.NewDefaultFactory(nameGenerator, colorGenerator)
	producer := message.NewProducer(clientManager)

	wsHandler := presentation.NewWebSocketHandler(clientFactory, clientManager, producer)

	http.HandleFunc("/", wsHandler.Handle)
	server := &http.Server{Addr: fmt.Sprintf(":%d", env.Server.Port)}

	gracefulShutdownCtx := graceful.Shutdown(&graceful.Params{
		OnStart:   func() { log.Printf("Graceful shutdown started. Waiting for active requests to complete") },
		OnTimeout: func() { log.Fatal("Graceful shutdown timed out. Forcing exit.") },
		OnShutdown: func(timeoutCtx context.Context) {
			if shutdownErr := server.Shutdown(timeoutCtx); shutdownErr != nil {
				log.Fatal(shutdownErr.Error())
			}
		},
	})

	log.Printf("Web server started listening on post %d", env.Server.Port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Could not initialize web server on port %d", env.Server.Port)
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
