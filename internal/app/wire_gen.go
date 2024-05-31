// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"cv/api/features"
	"cv/internal/app/router"
	"cv/internal/config"
	"cv/internal/infras/pgsql/repo"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

// Injectors from wire.go:

func InitApp(grpcServer *grpc.Server, log *slog.Logger) (*App, func(), error) {
	configConfig := config.New()
	dbEngine, cleanup, err := initDB(configConfig)
	if err != nil {
		return nil, nil, err
	}
	cvsRepository := repo.NewCvsPostgresRepo(dbEngine, log)
	featureClient, cleanup2, err := initFeaturesGRPC(configConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	cvsUsecases := usecases.NewCvsService(cvsRepository, featureClient)
	cvServiceServer := router.NewGRPCServer(grpcServer, cvsUsecases, log)
	app := NewApplication(configConfig, log, dbEngine, cvsUsecases, cvServiceServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func initDB(cfg *config.Config) (engine.DBEngine, func(), error) {
	connStr := engine.DBConnString(fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`, cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	db, err := engine.NewPostgresDB(connStr, 5, 2)
	if err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
}

func initFeaturesGRPC(cfg *config.Config) (features.FeatureClient, func(), error) {
	host := cfg.FeaturesClient.Host
	port := cfg.FeaturesClient.Port

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := features.NewFeatureClient(conn)
	return client, func() {
		conn.Close()
	}, nil
}
