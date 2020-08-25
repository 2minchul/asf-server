package service

import (
	"asf_server/config"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/hpcloud/tail"
	"net/http"
)

func GetLogPage(c *gin.Context) {
	cfg := c.MustGet("config").(config.Config)
	t, _ := tail.TailFile(cfg.AsfLogPath, tail.Config{Follow: false, MaxLineSize: 10})
	var b bytes.Buffer
	for line := range t.Lines {
		b.WriteString(line.Text)
		b.WriteByte('\n')
	}
	c.HTML(http.StatusOK, "log.html", gin.H{
		"text": b.String(),
	})
}
