package service

import (
	"asf_server/config"
	"github.com/gin-gonic/gin"
	"github.com/hpcloud/tail"
	"net/http"
	"strings"
)

func GetLogPage(c *gin.Context) {
	cfg := c.MustGet("config").(config.Config)
	t, _ := tail.TailFile(cfg.AsfLogPath, tail.Config{Follow: false})
	var lines = make([]string, 0, 10)
	for line := range t.Lines {
		lines = append(lines, line.Text)
	}
	if 10 < len(lines) {
		lines = lines[len(lines)-10:]
	}

	c.HTML(http.StatusOK, "log.html", gin.H{
		"text": strings.Join(lines, "\n"),
	})
}
