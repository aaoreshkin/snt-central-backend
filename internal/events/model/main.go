package model

import (
	"time"

	"github.com/oreshkindev/snt-central-backend/internal/attachments/model"
)

type (
	Event struct {
		ID          uint64             `db:"id" json:"id"`
		Title       string             `db:"title" json:"title"`
		Attachments []model.Attachment `db:"attachments" json:"attachments"`
		Published   bool               `db:"published" json:"published"`
		CreatedAt   time.Time          `db:"created_at" json:"created_at"`
		UpdatedAt   time.Time          `db:"updated_at" json:"updated_at"`
	}

	EventsAttachments struct {
		EventID      uint64    `json:"-"`
		AttachmentID uint64    `json:"-"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	Usecase interface {
		Create(*Event) (*Event, error)
		Find() ([]Event, error)
		First(uint64) (*Event, error)
		Update(*Event) (*Event, error)
		Delete(uint64) error
	}

	Repository interface {
		Create(*Event) (*Event, error)
		Find() ([]Event, error)
		First(uint64) (*Event, error)
		Update(*Event) (*Event, error)
		Delete(uint64) error
	}
)
