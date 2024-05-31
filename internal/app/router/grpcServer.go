package router

import (
	"context"
	"cv/api/cv"
	"cv/internal/usecases"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CvService struct {
	cv.UnimplementedCvServiceServer
	service usecases.CvsUsecases
	log     *slog.Logger
}

func (s *CvService) GetAll(context.Context, *cv.Paggination) (*cv.GetAllResponse, error) {
}

func (s *CvService) GetOne(context.Context, *cv.Id) (*cv.CV, error) {
	panic("unimplemented")
}

func (s *CvService) Upload(context.Context, *cv.CV) (*cv.Empty, error) {
	panic("unimplemented")
}

func NewGRPCServer(grpcServer *grpc.Server, service usecases.CvsUsecases, log *slog.Logger) cv.CvServiceServer {
	srv := &CvService{
		service: service,
		log:     log,
	}
	cv.RegisterCvServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	return srv
}
