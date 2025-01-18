package main

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB // Deklarasi variabel global db

func InitDatabase() {
    dsn := "root:@tcp(127.0.0.1:3306)/book_manager?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Gagal menghubungkan ke database: %v", err)
    }
    if err := db.AutoMigrate(&books{}); err != nil {
        log.Fatalf("Gagal migrasi schema: %v", err)
    }
    log.Println("Koneksi database berhasil!")
}

func GetDB() *gorm.DB {
    if db == nil {
        log.Fatal("Koneksi database belum diinisialisasi!")
    }
    return db
}
