package app

import (
	"context"
	"cv/api/cv"
	"cv/api/features"
	"cv/internal/config"
	"cv/internal/interceptors"
	"cv/internal/usecases"
	"cv/pkg/engine"
	"encoding/json"
	"log"
	"log/slog"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type App struct {
	Cfg  *config.Config
	Log  *slog.Logger
	pg   engine.DBEngine
	Nats *nats.Conn

	CvsService    usecases.CvsUsecases
	CvsGRPCServer cv.CvServiceServer

	js      nats.JetStreamContext
	fClient features.FeatureClient
}

func NewApplication(cfg *config.Config, log *slog.Logger, pg engine.DBEngine,
	cvsService usecases.CvsUsecases, grpcServer cv.CvServiceServer, nats *nats.Conn, featuresClient features.FeatureClient) *App {
	return &App{
		Cfg:           cfg,
		Log:           log,
		pg:            pg,
		CvsService:    cvsService,
		CvsGRPCServer: grpcServer,
		Nats:          nats,
		fClient:       featuresClient,
	}
}

func (a *App) Run() {
	js, _ := a.Nats.JetStream()
	sub, err := js.Subscribe("cv.feature", a.FeatureHandler, nil)
	if err != nil {
		log.Fatalf("Ошибка подписки на сообщения: %v", err)
	}
	defer sub.Unsubscribe()

	slog.Info("Server started")
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggingInterceptor((a.Log))))
	reflection.Register(s)
}

func (a *App) FeatureHandler(msg *nats.Msg) {
	type Message struct {
		CvID  string `json:"cvId"`
		Field string `json:"field"`
		Data  string `json:"data"`
	}

	// Распарсить JSON из сообщения
	var message Message
	err := json.Unmarshal(msg.Data, &message)
	if err != nil {
		log.Printf("Ошибка распарсивания JSON: %v", err)
		return
	}

	f, err := a.fClient.AddFeature(context.Background(), &features.FeatureStruct{
		Name:       message.Field,
		PriorityId: 0,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() != codes.AlreadyExists {
				return
			}
		}

		return
	}

	if err := a.CvsService.AddFeatureToCv(context.Background(), message.CvID, int(f.Id), message.Data); err != nil {
		slog.Error("cannot add feature to cv")
		return
	}
}
