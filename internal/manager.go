package internal

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/attachments"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/events"
	"github.com/oreshkindev/snt-central-backend/internal/invoices"
	"github.com/oreshkindev/snt-central-backend/internal/posts"
	"github.com/oreshkindev/snt-central-backend/internal/schedules"
	"github.com/oreshkindev/snt-central-backend/internal/user"
)

type Manager struct {
	User        *user.Manager
	Posts       *posts.Manager
	Event       *events.Manager
	Attachments *attachments.Manager
	Schedules   *schedules.Manager
	Invoices    *invoices.Manager
}

func New(context context.Context, connection *database.Connection) (*Manager, error) {
	return &Manager{
		User:        user.New(context, connection),
		Posts:       posts.New(context, connection),
		Event:       events.New(context, connection),
		Attachments: attachments.New(context, connection),
		Schedules:   schedules.New(context, connection),
		Invoices:    invoices.New(context, connection),
	}, nil
}
