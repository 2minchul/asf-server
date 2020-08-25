package middleware

import (
	"asf_server/config"
	"asf_server/token"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware(t token.Tokens) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("tokens", t)
		c.Next()
	}
}

func ConfigMiddleware(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	}
}
