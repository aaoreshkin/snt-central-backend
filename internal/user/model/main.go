package model

import "time"

type (
	User struct {
		ID           uint64    `db:"id" json:"id"`
		AccessToken  string    `db:"access_token" json:"access_token"`
		Description  string    `db:"description" json:"description"`
		Email        string    `db:"email" json:"email"`
		Fullname     string    `db:"fullname" json:"fullname"`
		Password     string    `db:"password" json:"password,omitempty"`
		PermissionID int64     `db:"permission_id" json:"permission_id"`
		Phone        string    `db:"phone" json:"phone"`
		Position     string    `db:"position" json:"position"`
		RefreshToken string    `db:"refresh_token" json:"refresh_token"`
		CreatedAt    time.Time `db:"created_at" json:"created_at"`
		UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	}

	Usecase interface {
		Create(*User) (*User, error)
		Find() ([]User, error)
		First(uint64) (*User, error)
		Update(*User) (*User, error)
		Delete(uint64) error
		Authenticate(*User) (*User, error)
		Revoke(string) (*User, error)
	}

	Repository interface {
		Create(*User) (*User, error)
		Find() ([]User, error)
		First(uint64) (*User, error)
		Any(string, string) (*User, error)
		Update(*User) (*User, error)
		Delete(uint64) error
	}
)
