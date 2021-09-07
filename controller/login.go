package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gurbanli/ProductManager/config"
	"github.com/gurbanli/ProductManager/dto"
	"github.com/gurbanli/ProductManager/service"
	"log"
	"net/http"
)

type LoginController struct{}

func (lController *LoginController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	ls := service.LoginService{}
	user, err := ls.LoginAccount(loginRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Email or password is incorrect"})
	} else {
		session, _ := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")
		if session != nil {
			session.Values["id"] = user.ID
			session.Values["authenticated"] = true
			err = session.Save(c.Request, c.Writer)
			if err != nil {
				log.Fatal("session error")
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "Logged In"})
	}
}
