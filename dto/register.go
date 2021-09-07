package dto

import "github.com/gurbanli/ProductManager/models"

type RegisterRequestDTO struct {
	Firstname       string            `json:"first_name" binding:"required"`
	LastName        string            `json:"last_name" binding:"required"`
	Email           string            `json:"email" binding:"required,email"`
	Password        string            `json:"password" binding:"required,min=8,max=32"`
	ConfirmPassword string            `json:"confirm_password" binding:"required,eqfield=Password"`
	BirthdayDate    models.CustomDate `json:"birthday_date" time_format:"2006-01-02" binding:"required"`
	IsAdmin         bool              `json:"is_admin"`
}
