package server

import (
	"golang-starter-data/config"
	"golang-starter-data/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			routers.RouteToRedis(v1.Group("/redis"))
			routers.RouteUser(v1.Group("/user"))
		}

	}

	os.Setenv("APP_ENV", "dev")
	cfg := config.GetConfig()

	err := r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
