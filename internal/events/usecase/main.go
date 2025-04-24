package usecase

import (
	"github.com/oreshkindev/snt-central-backend/internal/events/model"
)

type (
	Usecase struct {
		repository model.Repository
	}
)

func New(repository model.Repository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (usecase *Usecase) Create(entity *model.Event) (*model.Event, error) {
	return usecase.repository.Create(entity)
}

func (usecase *Usecase) Find() ([]model.Event, error) {
	return usecase.repository.Find()
}

func (usecase *Usecase) First(id uint64) (*model.Event, error) {
	return usecase.repository.First(id)
}

func (usecase *Usecase) Update(entity *model.Event) (*model.Event, error) {
	return usecase.repository.Update(entity)
}

func (usecase *Usecase) Delete(id uint64) error {
	return usecase.repository.Delete(id)
}
