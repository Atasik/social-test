package domain

import "errors"

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title" validate:"required"`
	Text  string `json:"text" validate:"required"`
}

type UpdatePostInput struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Text == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
