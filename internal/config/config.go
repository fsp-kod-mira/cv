package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		User string `env:"DB_USER" env-default:"postgres"`
		Pass string `env:"DB_PASS" env-default:"postgres"`
		Host string `env:"DB_HOST" env-default:"10.244.0.2"`
		Port string `env:"DB_PORT" env-default:"5436"`
		Name string `env:"DB_NAME" env-default:"cvs"`
	}
	Nats struct {
		Host string `env:"NATS_HOST" env-default:"46.29.236.28"`
		Port int    `env:"NATS_PORT" env-default:"4222"`
	}
	GRPC struct {
		Port string `env:"GRPC_PORT" env-default:"6001"`
	}

	FeaturesClient struct {
		Host string `env:"GRPC_FEATURES_HOST" env-default:"10.244.0.5"`
		Port string `env:"GRPC_FEATURES_PORT" env-default:"50052"`
	}
}

func New() *Config {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		header := "CV service env"
		f := cleanenv.FUsage(os.Stdout, cfg, &header)
		f()
		panic(err)
	}

	return cfg
}
