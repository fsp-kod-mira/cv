package usecases

import (
	"context"
	"cv/api/features"
	"cv/internal/domain"
	"fmt"
	"log/slog"

	"github.com/samber/lo"
)

type CvsService struct {
	CvsRepository
	featuresClient features.FeatureClient
}

func NewCvsService(cvsRepo CvsRepository, featuresClient features.FeatureClient) CvsUsecases {
	return &CvsService{
		CvsRepository:  cvsRepo,
		featuresClient: featuresClient,
	}
}

func (s *CvsService) GetOne(ctx context.Context, id string) (*domain.CV, error) {
	cv, err := s.CvsRepository.GetOne(ctx, id)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("error while get cv by id %d", err)
	}
	slog.Info("2")

	fs, err := s.GetFeaturesByCvId(ctx, id)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("error while get cv features by id %d", err)
	}

	slog.Info("3")
	domainCv := &domain.CV{
		Id:           id,
		Status:       cv.Status,
		FileId:       cv.FileId,
		UploadedById: cv.UploadedById,
		TotalScore:   cv.TotalScore,
	}

	slog.Info("4")
	for _, f := range fs {
		parsedFeature, err := s.featuresClient.GetFeaturesById(ctx, &features.IdStruct{
			Id: uint64(f.FeatureId),
		})
		if err != nil {
			continue
		}

		domainCv.Features = append(domainCv.Features, &domain.FullfieldFeature{
			FeatureName:  parsedFeature.FeatureName,
			PriorityName: parsedFeature.PriorityName,
			Coefficient:  parsedFeature.Coefficient,
		})
	}

	return domainCv, nil
}

func (s *CvsService) GetAll(ctx context.Context, limit, offset int) ([]*domain.CV, error) {
	cvsWithoutFields, err := s.CvsRepository.GetAll(ctx, limit, offset)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("errro while get all cvs")
	}

	return lo.Map(cvsWithoutFields, func(item *domain.CV, index int) *domain.CV {
		fs, err := s.GetFeaturesByCvId(ctx, item.Id)
		if err != nil {
			return nil
		}

		item.Features = make([]*domain.FullfieldFeature, 0)

		for _, f := range fs {
			parsedFeature, err := s.featuresClient.GetFeaturesById(ctx, &features.IdStruct{
				Id: uint64(f.FeatureId),
			})
			if err != nil {
				continue
			}

			item.Features = append(item.Features, &domain.FullfieldFeature{
				FeatureName:  parsedFeature.FeatureName,
				PriorityName: parsedFeature.PriorityName,
				Coefficient:  parsedFeature.Coefficient,
			})
		}

		return item
	}), nil
}
