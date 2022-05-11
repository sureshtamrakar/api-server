package routes

import (
	"github.com/gin-gonic/gin"
	controller_login "github.com/sureshtamrakar/api-server/controller/login"
	controller_register "github.com/sureshtamrakar/api-server/controller/register"
	controller_user "github.com/sureshtamrakar/api-server/controller/user"
)

func AddRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controller_login.Login)            // login user
	r.POST("/register", controller_register.CreateUser) // register user
	r.GET("/user", controller_user.Get)                 // get user
	return r
}
