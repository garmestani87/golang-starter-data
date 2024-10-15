package handlers

import (
	"golang-starter-data/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)


// @Summary put key and value
// @Description put key and value
// @Accept  json
// @Produce  json
// @Param   key     path    string     true     "key"
// @Param   value     path    string     true     "value"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/redis/{key}/{value} [put]
func Put(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")

	res := cache.GetClient().Set(key, value, 20*time.Second)
	if res.Err() != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, res.Err())
		return
	}

	ctx.JSON(http.StatusOK, "key/value put in to redis!")
}

// @Summary Show value for key
// @Description get value by key
// @Accept  json
// @Produce  json
// @Param   key     path    string     true     "key"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/redis/{key} [get]
func Get(ctx *gin.Context) {
	key := ctx.Param("key")

	var val *redis.StringCmd

	if val = cache.GetClient().Get(key); val.Err() != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, val.Err())
		return
	}

	ctx.JSON(http.StatusOK, val.String())
}
