package service

import (
	"context"
	"fmt"
	"log/slog"
)

type service struct{}

type IService interface {
	Greet(ctx context.Context, name string)
}

func InitService() IService {
	return &service{}
}

// Greet implements IService.
func (*service) Greet(ctx context.Context, name string) {
	slog.Info(fmt.Sprintf("Hello %s", name))
}
