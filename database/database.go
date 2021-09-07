package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Init(user string, password string, host string, port string, dbName string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Baku", host, port, user, password, dbName)
	err := error(nil)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can not connect to database !", err)
	}

}

func GetDB() *gorm.DB {
	return db
}
