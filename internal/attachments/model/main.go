package model

import (
	"mime/multipart"
	"time"
)

type (
	Attachment struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		Hex       string    `json:"hex"`
		Size      uint64    `json:"size"`
		Extension string    `json:"extension"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Usecase interface {
		Create(*multipart.FileHeader, *multipart.File) (*Attachment, error)
		Find() ([]*Attachment, error)
	}

	Repository interface {
		Create(*Attachment) (*Attachment, error)
		Find() ([]*Attachment, error)
	}
)
