package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type Config struct {
	Database `yaml:"database"`
	Server   `yaml:"server"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Server struct {
	Host        string        `yaml:"host"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

// MustLoad загружает конфигурацию из файла и переменных окружения
func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("config/local.yaml", &cfg); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	return &cfg
}
