package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetIp(c *gin.Context) {

	clientIP := c.GetHeader("X-Forwarded-For")
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	c.JSON(200, gin.H{"client-ip": clientIP})
}
