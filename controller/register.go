package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gurbanli/ProductManager/dto"
	"github.com/gurbanli/ProductManager/service"
	"net/http"
)

//author: gurbanli

type RegisterController struct{}

func (rController *RegisterController) Register(c *gin.Context) {

	var registerRequest dto.RegisterRequestDTO
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	rs := service.RegisterService{}
	user, err := rs.RegisterAccount(registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "registered", "user": user})

	}

}
