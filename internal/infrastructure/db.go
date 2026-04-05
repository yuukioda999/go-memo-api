package infrastructure

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB はデータベースの初期化と接続を行います
func InitDB() *gorm.DB {
	// test.db という名前でローカルファイルを生成/使用します
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connection established")
	return db
}
