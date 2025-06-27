package usecase

import (
	"github.com/oreshkindev/snt-central-backend/internal/invoices/model"
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

func (usecase *Usecase) Create(entity *model.Invoice) (*model.Invoice, error) {
	return usecase.repository.Create(entity)
}

func (usecase *Usecase) Find() ([]model.Invoice, error) {
	return usecase.repository.Find()
}

func (usecase *Usecase) First(id uint64) (*model.Invoice, error) {
	return usecase.repository.First(id)
}

func (usecase *Usecase) Update(entity *model.Invoice) (*model.Invoice, error) {
	return usecase.repository.Update(entity)
}

func (usecase *Usecase) Delete(id uint64) error {
	return usecase.repository.Delete(id)
}
