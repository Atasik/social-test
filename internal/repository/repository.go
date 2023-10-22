package repository

import (
	"social/internal/domain"
)

const (
	postsTable = "posts"
)

type PostRepo interface {
	Create(post domain.Post) (int, error)
	GetAll() ([]domain.Post, error)
	GetByID(id int) (domain.Post, error)
	Delete(id int) error
	Update(id int, input domain.UpdatePostInput) error
}

type Repository struct {
	PostRepo
}

func NewRepository(p PostRepo) *Repository {
	return &Repository{
		PostRepo: p,
	}
}
