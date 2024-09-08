package internal

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/post"
	"github.com/oreshkindev/snt-central-backend/internal/user"
)

type Manager struct {
	User *user.Manager
	Post *post.Manager
}

func New(context context.Context, connection *database.Connection) (*Manager, error) {

	return &Manager{
		User: user.New(context, connection),
		Post: post.New(context, connection),
	}, nil
}
