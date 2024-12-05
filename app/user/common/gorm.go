package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitGorm(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 连接成功后，可以进行数据库操作
	log.Println("Connected to PostgreSQL database")
	return db
}
