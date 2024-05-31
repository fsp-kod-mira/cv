package repo

import (
	"context"
	"cv/internal/domain"
	postgresql "cv/internal/infras/pgsql"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"log/slog"

	"github.com/google/uuid"
	lop "github.com/samber/lo/parallel"
)

type cvsPg struct {
	pg  engine.DBEngine
	log *slog.Logger
}

func NewCvsPostgresRepo(pg engine.DBEngine, log *slog.Logger) usecases.CvsRepository {
	return &cvsPg{pg: pg, log: log}
}

func (u *cvsPg) GetFeaturesByCvId(ctx context.Context, id string) ([]*domain.Feature, error) {
	querier := postgresql.New(u.pg.GetDB())
	features, err := querier.GetFeaturesByCvs(ctx, id)
	return lop.Map(features, func(item postgresql.CvsFeature, index int) *domain.Feature {
		return &domain.Feature{
			CvId:      id,
			FeatureId: int(item.FeatureID),
			Value:     item.Value,
		}
	}), err
}

func (u *cvsPg) Upload(ctx context.Context, cv *domain.CV) error {
	querier := postgresql.New(u.pg.GetDB())
	params := postgresql.CreateCvParams{
		ID:           uuid.New().String(),
		UploadedByID: cv.UploadedById,
		FileID:       cv.FileId,
	}

	if err := querier.CreateCv(ctx, params); err != nil {
		return CreationError("CV")
	}

	return nil
}

func (u *cvsPg) GetOne(ctx context.Context, id string) (*domain.CV, error) {
	querier := postgresql.New(u.pg.GetDB())
	cv, err := querier.GetCv(ctx, id)
	if err != nil {
		return nil, GetError("CV")
	}

	return &domain.CV{
		Id:           id,
		Status:       string(cv.Status),
		FileId:       cv.FileID,
		UploadedById: cv.UploadedByID,
	}, nil
}

// GetAll implements usecases.CvsRepository.
func (u *cvsPg) GetAll(ctx context.Context, limit int, offset int) ([]*domain.CV, error) {
	querier := postgresql.New(u.pg.GetDB())
	cvs, err := querier.GetAllCvs(ctx, postgresql.GetAllCvsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		slog.Error(err.Error())
		return nil, GetError("CV")
	}

	return lop.Map(cvs, func(item postgresql.Cv, index int) *domain.CV {
		return &domain.CV{
			Id:           item.ID,
			Status:       string(item.Status),
			FileId:       item.FileID,
			UploadedById: item.UploadedByID,
		}
	}), nil
}
