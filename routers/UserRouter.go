package routers

import (
	"github.com/gin-gonic/gin"
	"golang-starter-data/handlers"
)

func RouteUser(r *gin.RouterGroup) {
	r.POST("/", handlers.SaveUser)
	r.GET("/:id", handlers.FindUserById)
}
