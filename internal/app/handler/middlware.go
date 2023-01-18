package handler

import "github.com/gin-gonic/gin"

const xPingValue = "ping"

func (h *Handler) PingPong(c *gin.Context) {
	header := c.GetHeader("X-PING")

	if header == xPingValue {
		c.Header("X-PONG", "pong")
	}

	c.Next()
}
