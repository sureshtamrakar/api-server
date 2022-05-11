package controller_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_user "github.com/sureshtamrakar/api-server/models/user"
	util_auth "github.com/sureshtamrakar/api-server/util/auth"
)

func Get(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		return
	}
	id := c.GetInt("id")
	val, err := models_user.Load(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "id not found")
		return
	}
	c.JSON(http.StatusOK, val)
	return

}
