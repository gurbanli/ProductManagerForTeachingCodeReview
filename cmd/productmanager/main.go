package main

import (
	"github.com/gurbanli/ProductManager/config"
	"github.com/gurbanli/ProductManager/database"
	"github.com/gurbanli/ProductManager/models"
	"github.com/gurbanli/ProductManager/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	environment := os.Getenv("ENVIRONMENT")
	config.Init(environment)
	config.InitSessionConfig(environment == "production")

	conf := config.GetConfig()

	dbUsername := conf.GetString("database.username")
	dbPassword := conf.GetString("database.password")
	dbHost := conf.GetString("database.host")
	dbPort := conf.GetString("database.port")
	dbName := conf.GetString("database.name")

	database.Init(dbUsername, dbPassword, dbHost, dbPort, dbName)
	db := database.GetDB()

	err = db.AutoMigrate(models.User{}, models.Product{})
	if err != nil {
		log.Fatal("Can not migrate database !", err)
	}
	r := router.NewRouter(environment)
	err = r.Run(conf.GetString("host.address") + ":" + config.GetConfig().GetString("host.port"))
	if err != nil {
		log.Fatal(err)
	}
}
