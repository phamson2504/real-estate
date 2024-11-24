package config

import (
	"fmt"
	"log"
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	hostname = "localhost"
	port     = "3306"
	username = "root"
	password = "nhatban1"
	dbname   = "real_estate"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&model.User{}, &model.Property{}, &model.Image{},
		&model.Transaction{}, &model.Agent{}, &model.Favorite{},
		&model.Review{}, &model.City{}, &model.District{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	fmt.Println("Database connected and migrated successfully!")

	return db
}
