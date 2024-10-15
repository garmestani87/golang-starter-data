package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "golang-starter-data/docs"
)

func RouteSwagger(r *gin.RouterGroup) {
	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
