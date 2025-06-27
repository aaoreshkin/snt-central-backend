package model

import "time"

type (
	Invoice struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Published bool      `json:"published"`
		Priority  uint64    `json:"priority"`
		KwtPrice  float32   `json:"kwt_price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Usecase interface {
		Create(*Invoice) (*Invoice, error)
		Find() ([]Invoice, error)
		First(uint64) (*Invoice, error)
		Update(*Invoice) (*Invoice, error)
		Delete(uint64) error
	}

	Repository interface {
		Create(*Invoice) (*Invoice, error)
		Find() ([]Invoice, error)
		First(uint64) (*Invoice, error)
		Update(*Invoice) (*Invoice, error)
		Delete(uint64) error
	}
)
