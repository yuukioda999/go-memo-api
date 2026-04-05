package service

import (
	"errors"
	"go-memo-api/internal/domain"
	"go-memo-api/internal/repository"
)

type MemoService interface {
	CreateMemo(title, content string) (*domain.Memo, error)
	GetMemos() ([]domain.Memo, error)
}

type memoService struct {
	repo repository.MemoRepository
}

func NewMemoService(repo repository.MemoRepository) MemoService {
	return &memoService{repo: repo}
}

func (s *memoService) CreateMemo(title, content string) (*domain.Memo, error) {
	// バリデーション：タイトルが空なら保存させない
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	memo := &domain.Memo{
		Title:   title,
		Content: content,
	}

	if err := s.repo.Create(memo); err != nil {
		return nil, err
	}
	return memo, nil
}

func (s *memoService) GetMemos() ([]domain.Memo, error) {
	return s.repo.GetAll()
}
