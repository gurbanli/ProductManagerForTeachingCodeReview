package service

import (
	"github.com/gurbanli/ProductManager/database"
	"github.com/gurbanli/ProductManager/dto"
	"github.com/gurbanli/ProductManager/models"
	"github.com/gurbanli/ProductManager/utility"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type RegisterService struct {
}

func (rs *RegisterService) RegisterAccount(registerRequest dto.RegisterRequestDTO) (models.User, error) {
	user := models.User{
		FirstName:    registerRequest.Firstname,
		LastName:     registerRequest.LastName,
		Email:        registerRequest.Email,
		Password:     utility.GetHash(registerRequest.Password),
		BirthdayDate: time.Time(registerRequest.BirthdayDate),
		IsAdmin:      registerRequest.IsAdmin,
	}
	db = database.GetDB()
	result := db.Create(&user)
	return user, result.Error
}
