package repo

import (
	"context"
	"cv/internal/domain"
	postgresql "cv/internal/infras/pgsql"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"log/slog"

	"github.com/google/uuid"
)

type cvsPg struct {
	pg  engine.DBEngine
	log *slog.Logger
}

func NewCvsPostgresRepo(pg engine.DBEngine, log *slog.Logger) usecases.CvsRepository {
	return &cvsPg{pg: pg, log: log}
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
