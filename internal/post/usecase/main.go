package usecase

import "github.com/oreshkindev/snt-central-backend/internal/post/entity"

type PostUsecase struct {
	repository entity.PostRepository
}

func NewPostUsecase(repository entity.PostRepository) *PostUsecase {
	return &PostUsecase{
		repository: repository,
	}
}

func (usecase *PostUsecase) Create(entity *entity.Post) (*entity.Post, error) {
	return usecase.repository.Create(entity)
}

func (usecase *PostUsecase) Find() ([]entity.Post, error) {
	return usecase.repository.Find()
}

func (usecase *PostUsecase) First(id string) (*entity.Post, error) {
	return usecase.repository.First(id)
}

func (usecase *PostUsecase) Update(entity *entity.Post, id string) (*entity.Post, error) {
	return usecase.repository.Update(entity, id)
}

func (usecase *PostUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
