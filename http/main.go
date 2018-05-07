package http

import "github.com/gin-gonic/gin"

func defaultFunc(defaultResp interface{}) func(*gin.Context) {
	if defaultResp == nil {
		defaultResp = gin.H{
			"ping": "pong",
		}
	}
	return func(c *gin.Context) {
		c.JSON(200, defaultResp)
	}
}

// Run HTTP server
func Run() {
	r := gin.Default()
	// login
	r.POST("/login", defaultFunc(nil))
	// logout
	r.POST("/logout", defaultFunc(nil))
	// User Info
	r.GET("/:userId", defaultFunc(nil))

	r.Run()
}
