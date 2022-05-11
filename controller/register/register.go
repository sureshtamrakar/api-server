package controller_register

import (
	"net/http"

	models_login "github.com/sureshtamrakar/api-server/models/login"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var req models_login.Entity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	err := models_login.Create(req.Email, req.Name, req.Gender, req.Country, string(password), req.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to create user")
		return
	}
	c.JSON(http.StatusOK, "user created successfully")
}
