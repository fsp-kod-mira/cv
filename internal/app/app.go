package app

import (
	"cv/internal/config"
	"cv/internal/interceptors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() {
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggingInterceptor((a.log))))
	reflection.Register(s)
}
