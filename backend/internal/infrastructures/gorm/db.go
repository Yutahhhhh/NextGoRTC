package gorm

import (
	"os"
	"fmt"
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
	"NextGoRTC/internal/models"
)

func Connect() {
  USER := os.Getenv("DB_USER")
  PASS := os.Getenv("DB_PASSWORD")
  HOST := os.Getenv("DB_HOST")
  PORT := os.Getenv("DB_PORT")
  DBNAME := os.Getenv("DB_NAME")
  SSLMODE := "disable"
  dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tokyo", 
		HOST, USER, PASS, DBNAME, PORT, SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}

	log.Println("データベースに接続しました")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("データベースオブジェクトの取得に失敗しました: %v", err)
	}

	defer sqlDB.Close()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("自動マイグレーションに失敗しました: %v", err)
	}
}