package usecases

import (
	"context"
	"cv/internal/domain"
)

type CvsRepository interface {
	Upload(ctx context.Context, uploadedById string, fileId string) (string, error)
	GetOne(ctx context.Context, id string) (*domain.CV, error)
	GetAll(ctx context.Context, limit, offset int) ([]*domain.CV, error)
	GetFeaturesByCvId(ctx context.Context, id string) ([]*domain.Feature, error)
	AddFeatureToCv(ctx context.Context, cvId string, featureId int, value string) error
}

type CvsUsecases interface {
	Upload(ctx context.Context, uploadedById string, fileId string) (string, error)
	GetOne(ctx context.Context, id string) (*domain.CV, error)
	GetAll(ctx context.Context, limit, offset int) ([]*domain.CV, error)
	GetFeaturesByCvId(ctx context.Context, id string) ([]*domain.Feature, error)
	AddFeatureToCv(ctx context.Context, cvId string, featureId int, value string) error
}
