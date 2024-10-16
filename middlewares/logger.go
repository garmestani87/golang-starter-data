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
		console.Info("CLIENT_IP=", ctx.ClientIP(),
			", APP_NAME=", cfg.Logger.AppName,
			", TIME=", t.Format(time.RFC1123),
			", METHOD=", ctx.Request.Method,
			", URI=", ctx.Request.RequestURI,
			", STATUS=", status,
			", LATENCY=", latency,
			", USER_AGENT=", ctx.Request.UserAgent(),
		)
	}
}
