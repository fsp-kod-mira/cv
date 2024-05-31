package usecases

import (
	"context"
	"cv/internal/domain"
)

type CvsRepository interface {
	Upload(ctx context.Context, cv *domain.CV) error
	GetOne(ctx context.Context, id string) (*domain.CV, error)
	GetAll(ctx context.Context, limit, offset int) ([]*domain.CV, error)
	GetFeaturesByCvId(ctx context.Context, id string) ([]*domain.Feature, error)
}

type CvsUsecases interface {
	Upload(ctx context.Context, cv *domain.CV) error
	GetOne(ctx context.Context, id string) (*domain.CV, error)
	GetAll(ctx context.Context, limit, offset int) ([]*domain.CV, error)
	GetFeaturesByCvId(ctx context.Context, id string) ([]*domain.Feature, error)
}
