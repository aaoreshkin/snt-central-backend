package entity

import "time"

type (
	User struct {
		ID          uint64    `json:"id"`
		AccessToken string    `json:"access_token"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		Permission  string    `json:"permission"`
		Fullname    string    `json:"fullname"`
		Phone       string    `json:"phone"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	UserUsecase interface {
		Create(*User) (*User, error)
		Find() ([]User, error)
		First(string) (*User, error)
		Update(*User, uint64) (*User, error)
		Delete(uint64) error
		Authenticate(*User) (*User, error)
	}

	UserRepository interface {
		Create(*User) (*User, error)
		Find() ([]User, error)
		First(string) (*User, error)
		Update(*User, uint64) (*User, error)
		Delete(uint64) error
	}
)

// fields of struct that will be returned
func (response *User) NewResponse() *User {
	return &User{
		ID:          response.ID,
		AccessToken: response.AccessToken,
		Email:       response.Email,
		Password:    response.Password,
		Permission:  response.Permission,
		Fullname:    response.Fullname,
		Phone:       response.Phone,
		UpdatedAt:   response.UpdatedAt,
	}
}
