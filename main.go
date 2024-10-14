package main

import (
	"golang-starter-data/cache"
	"golang-starter-data/config"
	"golang-starter-data/server"
	"os"
)

func main() {
	setEnv()
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	server.Start(cfg)
}

func setEnv(){
	os.Setenv("APP_ENV", "dev")
}
