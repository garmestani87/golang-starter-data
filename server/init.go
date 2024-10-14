package server

import (
	"golang-starter-data/config"
	"golang-starter-data/routers"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	
	redis := r.Group("/redis")
	routers.RouteToRedis(redis)
	
	err := r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
