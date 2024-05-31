//go:build wireinject
// +build wireinject

package app

import (
	"cv/internal/app/router"
	"cv/internal/config"
	"cv/internal/infras/pgsql/repo"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"fmt"
	"log/slog"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitApp(grpcServer *grpc.Server, log *slog.Logger) (*App, func(), error) {
	panic(wire.Build(NewApplication,
		wire.NewSet(config.New),
		initDB,
		wire.NewSet(router.NewGRPCServer),
		wire.NewSet(repo.NewCvsPostgresRepo),
		wire.NewSet(usecases.NewCvsService),
	))
}

func initDB(cfg *config.Config) (engine.DBEngine, func(), error) {
	connStr := engine.DBConnString(fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`, cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	db, err := engine.NewPostgresDB(connStr, 5, 2)
	if err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
}
