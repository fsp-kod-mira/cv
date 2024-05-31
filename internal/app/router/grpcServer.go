package router

import (
	"context"
	"cv/api/cv"
	"cv/internal/domain"
	"cv/internal/usecases"
	"fmt"
	"log/slog"

	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CvService struct {
	cv.UnimplementedCvServiceServer
	service usecases.CvsUsecases
	log     *slog.Logger
}

// GetAll implements cv.CvServiceServer.
func (s *CvService) GetAll(ctx context.Context, pag *cv.Paggination) (*cv.GetAllResponse, error) {
	s.log.Info("trying to get cvs")
	cvs, err := s.service.GetAll(ctx, int(pag.Limit), int(pag.Offset))
	s.log.Info(fmt.Sprintf("successfully get %d cvs", len(cvs)))
	if err != nil {
		s.log.Error(err.Error())
		return nil, fmt.Errorf("error while get all cvs")
	}

	return &cv.GetAllResponse{
		Cvs: lo.Map(cvs, func(item *domain.CV, index int) *cv.CV {
			return &cv.CV{
				Id:           item.Id,
				Status:       item.Status,
				FileId:       item.FileId,
				UploadedById: item.UploadedById,
				Features: lo.Map(item.Features, func(fItem *domain.FullfieldFeature, index int) *cv.Feature {
					return &cv.Feature{
						FeatureName:  fItem.FeatureName,
						PriorityName: fItem.PriorityName,
						Coefficient:  fItem.Coefficient,
					}
				}),
			}
		}),
	}, nil
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
