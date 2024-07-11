package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT                string
	USER_SERVICE_PORT        string
	RESERVATION_SERVICE_PORT string
	PAYMENT_SERVICE_PORT     string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	cfg.USER_SERVICE_PORT = cast.ToString(coalesce("USER_SERVICE_PORT", ":8081"))
	cfg.RESERVATION_SERVICE_PORT = cast.ToString(coalesce("RESERVATION_SERVICE_PORT", ":8082"))
	cfg.PAYMENT_SERVICE_PORT = cast.ToString(coalesce("PAYMENT_SERVICE_PORT", ":8083"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
