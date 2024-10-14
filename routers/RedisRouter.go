package routers

import (
	"golang-starter-data/handlers"

	"github.com/gin-gonic/gin"
)

func RouteToRedis(r *gin.RouterGroup){
	r.PUT("/:key/:value", handlers.Put)
	r.GET("/:key", handlers.Get)
}