package server

import (
	"golang-starter-data/config"
	"golang-starter-data/middlewares"
	"golang-starter-data/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	cfg := config.GetConfig()
	
	r := gin.New()
	r.Use(middlewares.Logger(cfg), gin.Recovery())

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			routers.RouteToRedis(v1.Group("/redis"))
			routers.RouteUser(v1.Group("/user"))
			routers.RouteSwagger(v1.Group("/swagger"))
		}

	}

	os.Setenv("APP_ENV", "dev")

	err := r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
