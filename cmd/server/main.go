package main

import (
	"go-memo-api/internal/infrastructure" // 自分のモジュール名に合わせて調整
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. DBの初期化
	db := infrastructure.InitDB()

	// (後にここで repository などを初期化し、db を渡します)
	_ = db

	// 2. サーバーの構築
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}
