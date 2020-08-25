package service

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/hpcloud/tail"
	"net/http"
)

func GetLogPage(c *gin.Context) {
	t, _ := tail.TailFile("log.txt", tail.Config{Follow: false, MaxLineSize:10})
	var b bytes.Buffer
	for line := range t.Lines {
		b.WriteString(line.Text)
	}
	c.HTML(http.StatusOK, "log.html", gin.H{
		"text": b.String(),
	})
}