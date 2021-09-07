package models

import "time"

type User struct {
	Model
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email" gorm:"<-:create;unique"`
	Password     string    `json:"-"`
	BirthdayDate time.Time `json:"birthday_date"`
	IsAdmin      bool      `json:"-" gorm:"<-:create"`
	Products     []Product `json:"-"`
}
