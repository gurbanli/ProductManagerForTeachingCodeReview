package service

import (
	"github.com/gurbanli/ProductManager/database"
	"github.com/gurbanli/ProductManager/dto"
	"github.com/gurbanli/ProductManager/models"
	"github.com/gurbanli/ProductManager/utility"
)

type LoginService struct {
}

func (ls *LoginService) LoginAccount(loginRequest dto.LoginRequestDTO) (models.User, error) {
	user := models.User{}
	db := database.GetDB()
	result := db.Where("email = ? AND password = ?", loginRequest.Email, utility.GetHash(loginRequest.Password)).First(&user)
	return user, result.Error
}
