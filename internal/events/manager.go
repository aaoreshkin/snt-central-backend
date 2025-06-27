package events

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/events/controller"
	"github.com/oreshkindev/snt-central-backend/internal/events/repository"
	"github.com/oreshkindev/snt-central-backend/internal/events/usecase"
)

type Manager struct {
	Repository repository.Repository
	Usecase    usecase.Usecase
	Controller controller.Controller
}

func New(context context.Context, connection *database.Connection) *Manager {
	repository := repository.New(context, connection)
	usecase := usecase.New(repository)
	controller := controller.New(usecase)

	return &Manager{
		Repository: *repository,
		Usecase:    *usecase,
		Controller: *controller,
	}
}
