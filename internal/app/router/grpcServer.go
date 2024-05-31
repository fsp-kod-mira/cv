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

func (s *CvService) Upload(ctx context.Context, data *cv.UploadRequest) (*cv.Empty, error) {
	s.log.Info("trying to upload cv")
	if err := s.service.Upload(ctx, data.UploadedById, data.FileId); err != nil {
		s.log.Error(err.Error())
		return nil, fmt.Errorf("error while upload cv")
	}

	return nil, nil
}

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

func (s *CvService) GetOne(ctx context.Context, id *cv.Id) (*cv.CV, error) {
	s.log.Info("trying to get cv")
	if id == nil {
		return nil, fmt.Errorf("specify id")
	}

	item, err := s.service.GetOne(ctx, id.Id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, fmt.Errorf("error while get all cvs")
	}

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
	}, nil
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
