package repository

import (
	"errors"
	"social/internal/domain"
	"sync"
)

type postMemoryRepo struct {
	lastID int
	data   []domain.Post
	mutex  *sync.RWMutex
}

func NewPostMemoryRepo() *postMemoryRepo {
	return &postMemoryRepo{
		data:  []domain.Post{},
		mutex: &sync.RWMutex{},
	}
}

func (repo *postMemoryRepo) Create(post domain.Post) (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.lastID++
	post.ID = repo.lastID
	repo.data = append(repo.data, post)
	return repo.lastID, nil

}

func (repo *postMemoryRepo) GetAll() ([]domain.Post, error) {
	return repo.data, nil
}

func (repo *postMemoryRepo) GetByID(id int) (domain.Post, error) {
	for _, post := range repo.data {
		if post.ID == id {
			return post, nil
		}
	}
	return domain.Post{}, errors.New("not found")
}

func (repo *postMemoryRepo) Delete(id int) error {
	i := -1
	for idx, post := range repo.data {
		if post.ID != id {
			continue
		}
		i = idx
	}

	if i < 0 {
		return errors.New("not found")
	}

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	if i < len(repo.data)-1 {
		copy(repo.data[i:], repo.data[i+1:])
	}
	repo.data[len(repo.data)-1] = domain.Post{}
	repo.data = repo.data[:len(repo.data)-1]
	return nil
}

func (repo *postMemoryRepo) Update(id int, input domain.UpdatePostInput) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	for _, post := range repo.data {
		if post.ID == id {
			if input.Text != nil {
				post.Text = *input.Text
			}
			if input.Title != nil {
				post.Title = *input.Title
			}
		}
	}
	return errors.New("not found")
}
