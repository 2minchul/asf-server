package service

import (
	"asf_server/asf_cmd"
	"asf_server/token"
	"github.com/gin-gonic/gin"
	"log"
)

func getTokens(c *gin.Context) token.Tokens {
	return c.MustGet("tokens").(token.Tokens)
}

func checkAvailableToken(c *gin.Context) bool {
	tokens := getTokens(c)

	token.RemoveExpiredToken(tokens)
	tokenStr := c.PostForm("token")
	_, isExist := tokens[tokenStr]
	if isExist {
		delete(tokens, tokenStr)
	}
	return isExist
}

func CreateToken(c *gin.Context) {
	tokens := getTokens(c)

	token.RemoveExpiredToken(tokens)
	tokenStr := token.CreateToken(tokens)
	log.Printf("%v\n", len(tokens))
	c.JSON(200, gin.H{
		"token": tokenStr,
	})
}

func RunAsf(c *gin.Context) {
	if !checkAvailableToken(c) {
		c.JSON(403, gin.H{"status": "fail"})
		return
	}

	asf_cmd.StartAsf()
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func StopAsf(c *gin.Context) {
	if !checkAvailableToken(c) {
		c.JSON(403, gin.H{"status": "fail"})
		return
	}

	asf_cmd.StopAsf()
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
