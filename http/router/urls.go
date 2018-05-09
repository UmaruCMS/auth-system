package router

import (
	"github.com/UmaruCMS/auth-system/http/api"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	api.RegisterUserHandlers(r)
}
