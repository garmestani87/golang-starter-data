package main

import (
	"golang-starter-data/cache"
	"golang-starter-data/config"
	"golang-starter-data/db"
	"golang-starter-data/server"
	"os"
)

func main() {
	setEnv()
	cfg := config.GetConfig()
	initRedis(cfg)
	initPostgres(cfg)
	server.Start()
}

func setEnv() {
	os.Setenv("APP_ENV", "dev")
}

func initRedis(cfg *config.Config) {
	// defer cache.CloseRedis()

	cache.InitRedis(cfg)
}

func initPostgres(cfg *config.Config) {
	// defer db.CloseDatabase()

	db.InitPostgres(cfg)
}
