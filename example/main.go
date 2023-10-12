package main

import (
	"context"
	"log/slog"

	"github.com/swagftw/gi"
	"github.com/swagftw/gi/example/service"
)

func main() {
	svc := service.InitService()

	// Inject the service
	err := gi.Inject(&svc)
	if err != nil {
		slog.Error("failed to inject service", "error", err)

		return
	}

	// Invoke service from anywhere in codebase
	invokedSvc, err := gi.Invoke[*service.IService]()
	if err != nil {
		slog.Error("failed to invoke service", "error", err)

		return
	}

	(*invokedSvc).Greet(context.Background(), "Sean") // prints "Hello Sean" to standard output
}
