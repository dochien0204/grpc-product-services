package db

import (
	"fmt"
	"log"
	"product_svc/pkg/config"
	"product_svc/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func ConnectDB(c config.DBConfig) Handler {
	host := c.Host
	user := c.User
	password := c.Password
	databaseName := c.DatabaseName
	port := c.Port

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect DB")
	}

	fmt.Println("Connect DB successfully")

	db.AutoMigrate(&models.Product{})

	return Handler{db}
}
