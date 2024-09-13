package usecase

import (
	"github.com/oreshkindev/rbac-middleware"
	"github.com/oreshkindev/snt-central-backend/common"
	"github.com/oreshkindev/snt-central-backend/internal/user/entity"
)

type (
	UserUsecase struct {
		repository entity.UserRepository
	}

	Subject struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}
)

func NewUserUsecase(repository entity.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (usecase *UserUsecase) Create(entity *entity.User) (*entity.User, error) {

	// Hash entity raw password
	hashedPassword, err := common.HashPassword(entity.Password)
	if err != nil {
		return nil, err
	}

	// Set entity hashed password
	entity.Password = hashedPassword

	hashedToken, err := rbac.HashToken(Subject{Email: entity.Email, Role: entity.Permission}, 15)
	if err != nil {
		return nil, err
	}

	// Set entity access token
	entity.AccessToken = hashedToken

	return usecase.repository.Create(entity)
}

func (usecase *UserUsecase) Find() ([]entity.User, error) {
	return usecase.repository.Find()
}

func (usecase *UserUsecase) First(email string) (*entity.User, error) {
	return usecase.repository.First(email)
}

func (usecase *UserUsecase) Update(entity *entity.User, id uint64) (*entity.User, error) {

	// Hash entity raw password
	hashedPassword, err := common.HashPassword(entity.Password)
	if err != nil {
		return nil, err
	}

	// Set entity hashed password
	entity.Password = hashedPassword

	return usecase.repository.Update(entity, id)
}

func (usecase *UserUsecase) Delete(id uint64) error {
	return usecase.repository.Delete(id)
}

func (usecase *UserUsecase) Authenticate(entity *entity.User) (*entity.User, error) {

	// Check is user exist
	exist, err := usecase.repository.First(entity.Email)
	if err != nil {
		return nil, err
	}

	// Check is password equal to exist
	equal, err := common.CheckPasswordHash(entity.Password, exist.Password)
	if err != nil {
		return nil, err
	}

	if !equal {
		return nil, err
	}

	hashedToken, err := rbac.HashToken(Subject{Email: exist.Email, Role: exist.Permission}, 15)
	if err != nil {
		return nil, err
	}

	// Set access token
	exist.AccessToken = hashedToken

	result, err := usecase.repository.Update(exist, exist.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
