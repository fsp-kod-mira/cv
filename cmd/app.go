package main

import (
	"cv/internal/app"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	a, cleanup, err := app.InitApp(srv, log)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	log.Debug("init application")
	l, err := net.Listen("tcp", ":"+a.Cfg.GRPC.Port)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	defer func() {
		if err1 := l.Close(); err != nil {
			log.Error(err.Error())
			panic(err1)
		}
	}()

	log.Debug("grpc server start")
	if err := srv.Serve(l); err != nil {
		log.Error(err.Error())
		panic(err)
	}

	log.Debug("waiting for shutdown signal")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		cleanup()
	}
}
