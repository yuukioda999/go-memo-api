package repository

import (
	"go-memo-api/internal/domain"

	"gorm.io/gorm"
)

// MemoRepository は、外部から呼び出すための「窓口（インターフェース）」です
type MemoRepository interface {
	Create(memo *domain.Memo) error
	GetAll() ([]domain.Memo, error)
}

// memoRepository は、実際に GORM (DB) を使って動く「実体」です
type memoRepository struct {
	db *gorm.DB
}

// NewMemoRepository は、このリポジトリを使えるように準備する関数です
func NewMemoRepository(db *gorm.DB) MemoRepository {
	return &memoRepository{db: db}
}

// Create は、新しいメモを1件保存します
func (r *memoRepository) Create(memo *domain.Memo) error {
	return r.db.Create(memo).Error
}

// GetAll は、保存されているすべてのメモを取得します
func (r *memoRepository) GetAll() ([]domain.Memo, error) {
	var memos []domain.Memo
	// IDの降順（新しい順）で取得する設定にしています
	err := r.db.Order("id desc").Find(&memos).Error
	return memos, err
}
