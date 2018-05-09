package api

import (
	"net/http"
	"strconv"

	"github.com/UmaruCMS/auth-system/controller/user"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	c.String(http.StatusNotFound, "url not found")
}

type loginForm struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func login(c *gin.Context) {
	loginForm := &loginForm{}
	err := c.ShouldBind(loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, token := user.Login(loginForm.Email, loginForm.Password)
	if ok {
		c.String(http.StatusOK, token)
	} else {
		c.String(http.StatusUnauthorized, "Authorization failed")
	}

}

func getUserInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Not a vaild integer ID")
		return
	}

	user := user.GetUserInfo(uint(userID))

	if user == nil {
		c.String(http.StatusNotFound, "Not found user %d", userID)
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func RegisterUserHandlers(r *gin.Engine) {
	ur := r.Group("/users")
	ur.POST("/register", register)
	ur.POST("/login", login)
	ur.GET("/:userID", getUserInfo)
}
