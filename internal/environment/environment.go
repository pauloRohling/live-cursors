package environment

import "time"

type Config struct {
	Server struct {
		Port int `yml:"port" env:"SERVER_PORT"`
	} `yml:"server"`
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
