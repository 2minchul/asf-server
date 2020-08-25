package main

import (
	"asf_server/config"
	"asf_server/middleware"
	"asf_server/service"
	"asf_server/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var tokens token.Tokens

func main() {
	tokens = make(token.Tokens)
	r := gin.Default()
	r.Use(middleware.ConfigMiddleware(config.GetConfig("config.ini")))
	r.Use(middleware.TokenMiddleware(tokens))

	r.LoadHTMLGlob("templates/*")
	r.StaticFS("/static/highlight", http.Dir("static/highlight"))
	r.StaticFile("/static/axios.min.js", "static/axios.min.js")

	// api
	r.POST("/api/token", service.CreateToken)
	r.POST("/api/run", service.RunAsf)
	r.POST("/api/stop", service.StopAsf)

	// html
	r.StaticFile("/start", "static/start.html")
	r.StaticFile("/stop", "static/stop.html")
	r.GET("/log", service.GetLogPage)

	//// debug
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Main website",
	//	})
	//})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(r.Run("0.0.0.0:1022"))
}
