package user

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/user/controller"
	"github.com/oreshkindev/snt-central-backend/internal/user/repository"
	"github.com/oreshkindev/snt-central-backend/internal/user/usecase"
)

type Manager struct {
	UserRepository repository.UserRepository
	UserUsecase    usecase.UserUsecase
	UserController controller.UserController
}

func New(context context.Context, connection *database.Connection) *Manager {
	userRepository := repository.NewUserRepository(context, connection)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	return &Manager{
		UserRepository: *userRepository,
		UserUsecase:    *userUsecase,
		UserController: *userController,
	}
}
