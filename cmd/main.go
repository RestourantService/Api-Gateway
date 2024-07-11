package main

import (
	"api-gateway/api"
	"api-gateway/config"
)

func main() {
	cfg := config.Load()

	router := api.NewRouter(cfg)
	router.Run(cfg.HTTP_PORT)
}
