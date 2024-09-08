package post

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/post/controller"
	"github.com/oreshkindev/snt-central-backend/internal/post/repository"
	"github.com/oreshkindev/snt-central-backend/internal/post/usecase"
)

type Manager struct {
	PostRepository repository.PostRepository
	PostUsecase    usecase.PostUsecase
	PostController controller.PostController
}

func New(context context.Context, connection *database.Connection) *Manager {
	postRepository := repository.NewPostRepository(context, connection)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postController := controller.NewPostController(postUsecase)

	return &Manager{
		PostRepository: *postRepository,
		PostUsecase:    *postUsecase,
		PostController: *postController,
	}
}
