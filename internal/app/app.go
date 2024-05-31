package app

import (
	"cv/api/cv"
	"cv/internal/config"
	"cv/internal/interceptors"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	Cfg *config.Config
	Log *slog.Logger
	pg  engine.DBEngine

	CvsService    usecases.CvsUsecases
	CvsGRPCServer cv.CvServiceServer
}

func NewApplication(cfg *config.Config, log *slog.Logger, pg engine.DBEngine,
	cvsService usecases.CvsUsecases, grpcServer cv.CvServiceServer) *App {
	return &App{
		Cfg:           cfg,
		Log:           log,
		pg:            pg,
		CvsService:    cvsService,
		CvsGRPCServer: grpcServer,
	}
}

func (a *App) Run() {
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggingInterceptor((a.Log))))
	reflection.Register(s)
}
