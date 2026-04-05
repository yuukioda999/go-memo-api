package domain

import (
	"time"
)

// Memo はこのアプリケーションで扱う「メモ」の構造体です
type Memo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:100;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
