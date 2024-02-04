package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env 		string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer	`yaml:"http_server"`
}

type HTTPServer struct {
	Address		 string			`yaml:"address" env-default:"localhost:8080"`
	Timeout		 time.Duration	`yaml:"timeout" env-default:"4s"`
	IdleTimeout  time.Duration 	`yaml:"idle_timeout" env-default:"60s"`
	User		 string			`yaml:"user" env-required:"true"`
	Password	 string			`yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

// Must-функции не возвращают ошибку, а паникуют
func MustLoad() *Config {
	configPath := "config/local.yaml"	// переписать через os.Getenv
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	// reading configuration file
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}