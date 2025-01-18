package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func initDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/book_manager?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrasi model ke database
	db.AutoMigrate(&Book{})
}
