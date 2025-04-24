package usecase

import (
	"github.com/oreshkindev/rbac-middleware"
	"github.com/oreshkindev/snt-central-backend/common"
	"github.com/oreshkindev/snt-central-backend/internal/user/model"
)

type (
	Usecase struct {
		repository model.Repository
	}
)

const (
	timeout         = 15
	timeoutDuration = 480
)

func New(repository model.Repository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (usecase *Usecase) Create(entity *model.User) (*model.User, error) {
	var err error

	// Hash entity raw password
	if entity.Password, err = common.HashPassword(entity.Password); err != nil {
		return nil, err
	}

	if entity.AccessToken, err = rbac.Hash(map[string]interface{}{
		"email":      entity.Email,
		"permission": entity.PermissionID,
	}, timeout); err != nil {
		return nil, err
	}

	if entity.RefreshToken, err = rbac.Hash(map[string]interface{}{}, timeoutDuration); err != nil {
		return nil, err
	}

	return usecase.repository.Create(entity)
}

func (usecase *Usecase) Find() ([]model.User, error) {
	return usecase.repository.Find()
}

func (usecase *Usecase) First(id uint64) (*model.User, error) {
	return usecase.repository.First(id)
}

func (usecase *Usecase) Update(entity *model.User) (*model.User, error) {
	var err error

	// Получаем текущего пользователя из базы данных
	exist, err := usecase.repository.Any("phone", entity.Phone)
	if err != nil {
		return nil, err
	}

	// Если пароль не пустой, хешируем его, иначе используем старый пароль
	if entity.Password != "" {
		if entity.Password, err = common.HashPassword(entity.Password); err != nil {
			return nil, err
		}
	} else {
		entity.Password = exist.Password
	}

	if entity.AccessToken, err = rbac.Hash(map[string]any{
		"email":      exist.Email,
		"permission": exist.PermissionID,
	}, timeout); err != nil {
		return nil, err
	}

	if entity.RefreshToken, err = rbac.Hash(map[string]any{}, timeoutDuration); err != nil {
		return nil, err
	}

	return usecase.repository.Update(entity)
}

func (usecase *Usecase) Delete(id uint64) error {
	return usecase.repository.Delete(id)
}

func (usecase *Usecase) Authenticate(entity *model.User) (*model.User, error) {

	var err error

	// Check is user exist
	exist, err := usecase.repository.Any("phone", entity.Phone)
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

	if exist.AccessToken, err = rbac.Hash(map[string]any{
		"email":      exist.Email,
		"permission": exist.PermissionID,
	}, timeout); err != nil {
		return nil, err
	}

	if exist.RefreshToken, err = rbac.Hash(map[string]any{}, timeoutDuration); err != nil {
		return nil, err
	}

	return usecase.repository.Update(exist)
}

func (usecase *Usecase) Revoke(refresh_token string) (*model.User, error) {

	var err error

	if _, err := rbac.Validate(refresh_token); err != nil {
		return nil, err
	}

	// Check is user exist
	exist, err := usecase.repository.Any("refresh_token", refresh_token)
	if err != nil {
		return nil, err
	}

	if exist.AccessToken, err = rbac.Hash(map[string]any{
		"email":      exist.Email,
		"permission": exist.PermissionID,
	}, timeout); err != nil {
		return nil, err
	}

	if exist.RefreshToken, err = rbac.Hash(map[string]any{}, timeoutDuration); err != nil {
		return nil, err
	}

	return usecase.repository.Update(exist)
}
