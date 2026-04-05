package main

import (
	"go-memo-api/internal/handler"
	"go-memo-api/internal/infrastructure"
	"go-memo-api/internal/repository"
	"go-memo-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. インフラ（DB）
	db := infrastructure.InitDB()

	// 2. 依存関係の注入 (Dependency Injection)
	repo := repository.NewMemoRepository(db)
	svc := service.NewMemoService(repo)
	h := handler.NewMemoHandler(svc)

	// 3. ルーティング
	r := gin.Default()

	r.POST("/memos", h.Create)
	r.GET("/memos", h.GetAll)

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
