package handlers

import (
	"golang-starter-data/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Put(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")

	res := cache.GetClient().Set(key, value, 20*time.Second)
	if res.Err() != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Err())
		return
	}

	ctx.JSON(http.StatusOK, "key/value put in to redis!")
}

func Get(ctx *gin.Context) {
	key := ctx.Param("key")

	var val *redis.StringCmd

	if val = cache.GetClient().Get(key); val.Err() != nil {
		ctx.AbortWithError(http.StatusInternalServerError, val.Err())
		return
	}

	ctx.JSON(http.StatusOK, val.String())
}
