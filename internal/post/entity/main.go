package entity

import "time"

type (
	Post struct {
		ID          uint64    `json:"id"`
		Promo       bool      `json:"promo"`
		Published   bool      `json:"published"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
	}

	PostUsecase interface {
		Create(*Post) (*Post, error)
		Find() ([]Post, error)
		First(string) (*Post, error)
		Update(*Post, string) (*Post, error)
		Delete(string) error
	}

	PostRepository interface {
		Create(*Post) (*Post, error)
		Find() ([]Post, error)
		First(string) (*Post, error)
		Update(*Post, string) (*Post, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Post) NewResponse() *Post {
	return &Post{
		ID:          response.ID,
		Promo:       response.Promo,
		Published:   response.Published,
		Title:       response.Title,
		Description: response.Description,
		CreatedAt:   response.CreatedAt,
	}
}
