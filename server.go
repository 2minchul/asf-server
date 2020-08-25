package main

import (
	"asf_server/service"
	"asf_server/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var tokens token.Tokens

func TokenMiddleware(t token.Tokens) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("tokens", t)
		c.Next()
	}
}

func main() {
	tokens = make(token.Tokens)
	r := gin.Default()
	r.Use(TokenMiddleware(tokens))

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

	log.Println("http://127.0.0.1:8080")
	log.Fatal(r.Run("127.0.0.1:8080"))
}
