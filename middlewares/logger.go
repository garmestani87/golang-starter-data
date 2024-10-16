package middlewares

import (
	"golang-starter-data/config"
	"golang-starter-data/loggers"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()
		latency := time.Since(t)

		status := ctx.Writer.Status()

		console := loggers.NewConsoleLogger()
		console.Info(ctx.ClientIP(),
			cfg.Logger.AppName,
			t.Format(time.RFC1123),
			ctx.Request.Method,
			ctx.Request.RequestURI,
			status,
			latency,
			ctx.Request.UserAgent(),
		)
	}
}
