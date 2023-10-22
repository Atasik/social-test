package service

import (
	"social/internal/domain"
	"social/internal/repository"
)

type postService struct {
	postRepo repository.PostRepo
}

func NewPostService(postRepo repository.PostRepo) *postService {
	return &postService{postRepo: postRepo}
}

func (s *postService) Create(post domain.Post) (int, error) {
	return s.postRepo.Create(post)
}

func (s *postService) GetAll() ([]domain.Post, error) {
	return s.postRepo.GetAll()
}

func (s *postService) GetByID(id int) (domain.Post, error) {
	return s.postRepo.GetByID(id)
}

func (s *postService) Delete(id int) error {
	return s.postRepo.Delete(id)
}

func (s *postService) Update(id int, input domain.UpdatePostInput) error {
	return s.postRepo.Update(id, input)
}
