package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		User string `env:"DB_USER" env-default:"postgres"`
		Pass string `env:"DB_PASS" env-default:"postgres"`
		Host string `env:"DB_HOST" env-default:"localhost"`
		Port int    `env:"DB_PORT" env-default:"5436"`
		Name string `env:"DB_NAME" env-default:"users"`
	}
	GRPC struct {
		Port string `env:"GRPC_PORT" env-default:"6001"`
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
