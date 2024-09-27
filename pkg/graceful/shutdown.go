package graceful

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Params struct {
	OnStart    func()
	OnTimeout  func()
	OnShutdown func(timeoutCtx context.Context)
}

func Shutdown(params *Params) context.Context {
	if params == nil {
		params = &Params{}
	}

	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	applicationCtx, applicationStopCtx := context.WithCancel(context.Background())

	go func() {
		<-interruptSignal

		if params.OnStart != nil {
			params.OnStart()
		}

		timeoutCtx, _ := context.WithTimeout(applicationCtx, 30*time.Second)

		go func() {
			<-timeoutCtx.Done()
			if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) && params.OnTimeout != nil {
				params.OnTimeout()
			}
		}()

		if params.OnShutdown != nil {
			params.OnShutdown(timeoutCtx)
		}
		applicationStopCtx()
	}()

	return applicationCtx
}
