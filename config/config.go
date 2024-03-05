package config

import (
	"os"
	"strings"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	OpenDotaAPI struct {
		KeyPath string `env:"OPEN_DOTA_API_KEY_PATH" env-required:"true"`
		Key     string
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"8089"`
		Host string `env:"HTTP_HOST" env-default:"0.0.0.0"`
	}
	Postgres struct {
		Port         string `env:"DB_PORT" env-default:"5432"`
		Host         string `env:"DB_HOST" env-default:"postgres"`
		Username     string `env:"DB_USERNAME" env-default:"postgres"`
		DatabaseName string `env:"DB_NAME" env-default:"main"`
		PasswordPath string `env:"DB_PASSWORD_PATH" env-required:"true"`
		Password     string
	}
	Redis struct {
		Addr string `env:"REDIS_ADDR" env-default:"redis:6379"`
	}
	FrontendDir string `env:"FRONTEND_DIR" env-default:"./dist"`
}

var cfg Config
var once sync.Once

func InitConfig() Config {
	once.Do(func() {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			panic(err)
		}

		apiKey, err := os.ReadFile(cfg.OpenDotaAPI.KeyPath)
		if err != nil {
			panic(err)
		}

		dbPassword, err := os.ReadFile(cfg.Postgres.PasswordPath)
		if err != nil {
			panic(err)
		}

		cfg.OpenDotaAPI.Key = strings.TrimSpace(string(apiKey))
		cfg.Postgres.Password = strings.TrimSpace(string(dbPassword))
	})
	return cfg
}
