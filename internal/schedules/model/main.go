package model

import "time"

type (
	Schedule struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Published bool      `json:"published"`
		Priority  uint64    `json:"priority"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Usecase interface {
		Create(*Schedule) (*Schedule, error)
		Find() ([]Schedule, error)
		First(uint64) (*Schedule, error)
		Update(*Schedule) (*Schedule, error)
		Delete(uint64) error
	}

	Repository interface {
		Create(*Schedule) (*Schedule, error)
		Find() ([]Schedule, error)
		First(uint64) (*Schedule, error)
		Update(*Schedule) (*Schedule, error)
		Delete(uint64) error
	}
)
