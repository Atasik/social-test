package service

import (
	"social/internal/domain"
)

type Post interface {
	Create(post domain.Post) (int, error)
	GetAll() ([]domain.Post, error)
	GetByID(id int) (domain.Post, error)
	Delete(id int) error
	Update(id int, input domain.UpdatePostInput) error
}

type Service struct {
	Post
}

func NewService(p Post) *Service {
	return &Service{
		Post: p,
	}
}
