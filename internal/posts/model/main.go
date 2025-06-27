package model

import "time"

type (
	Post struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Published bool      `json:"published"`
		Ads       bool      `json:"ads"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Usecase interface {
		Create(*Post) (*Post, error)
		Find() ([]Post, error)
		First(uint64) (*Post, error)
		Update(*Post) (*Post, error)
		Delete(uint64) error
	}

	Repository interface {
		Create(*Post) (*Post, error)
		Find() ([]Post, error)
		First(uint64) (*Post, error)
		Update(*Post) (*Post, error)
		Delete(uint64) error
	}
)
