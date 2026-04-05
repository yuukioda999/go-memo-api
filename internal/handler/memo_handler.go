package handler

import (
	"go-memo-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemoHandler struct {
	svc service.MemoService
}

func NewMemoHandler(svc service.MemoService) *MemoHandler {
	return &MemoHandler{svc: svc}
}

// Create は POST /memos を処理します
func (h *MemoHandler) Create(c *gin.Context) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memo, err := h.svc.CreateMemo(input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, memo)
}

// GetAll は GET /memos を処理します
func (h *MemoHandler) GetAll(c *gin.Context) {
	memos, err := h.svc.GetMemos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memos)
}
