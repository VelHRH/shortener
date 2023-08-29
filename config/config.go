package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string        `yaml:"address" env-default:"localhost:4444"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
}

func MustLoad() *Config {
	cp := "C:/xampp/htdocs/GOLANG/shortener/config/local.yaml"
	if _, err := os.Stat(cp); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", cp)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cp, &cfg); err != nil {
		log.Fatalf("failed to read config: %s", cp)
	}

	return &cfg
}
