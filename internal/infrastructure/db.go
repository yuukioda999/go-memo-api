package infrastructure

import (
	"go-memo-api/internal/domain" // ここを追加（自分のモジュール名に合わせてください）
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// ★ ここを追加：構造体の定義に合わせてテーブルを作成・更新します
	err = db.AutoMigrate(&domain.Memo{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database connection established and migration completed")
	return db
}
