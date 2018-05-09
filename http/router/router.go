package router

import "github.com/gin-gonic/gin"

func DefaultRouter() *gin.Engine {
	r := gin.Default()
	return r
}
