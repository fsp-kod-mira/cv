//go:build wireinject
// +build wireinject

package app

import (
	"cv/api/features"
	"cv/internal/app/router"
	"cv/internal/config"
	"cv/internal/infras/pgsql/repo"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"fmt"
	"log/slog"

	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitApp(grpcServer *grpc.Server, log *slog.Logger) (*App, func(), error) {
	panic(wire.Build(NewApplication,
		wire.NewSet(config.New),
		initDB,
		initFeaturesGRPC,
		initNats,
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

func initFeaturesGRPC(cfg *config.Config) (features.FeatureClient, func(), error) {
	host := cfg.FeaturesClient.Host
	port := cfg.FeaturesClient.Port

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	client := features.NewFeatureClient(conn)
	return client, func() {
		conn.Close()
	}, nil
}

func initNats(cfg *config.Config) (*nats.Conn, func(), error) {
	nc, err := nats.Connect(fmt.Sprintf("nats://%s:%d", cfg.Nats.Host, cfg.Nats.Port))
	if err != nil {
		return nil, nil, err
	}
	return nc, func() {
		nc.Close()
	}, nil
}
