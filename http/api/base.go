package api

import "github.com/gin-gonic/gin"

func DefaultHandlerFactory(defaultResp interface{}) func(*gin.Context) {
	if defaultResp == nil {
		defaultResp = gin.H{
			"ping": "pong",
		}
	}
	return func(c *gin.Context) {
		c.JSON(200, defaultResp)
	}
}
