package repository

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/schedules/model"
)

type (
	Repository struct {
		context    context.Context
		connection *database.Connection
	}
)

func New(context context.Context, connection *database.Connection) *Repository {
	return &Repository{context, connection}
}

func (repository *Repository) Create(entity *model.Schedule) (*model.Schedule, error) {
	const query = `
	INSERT INTO schedules (
		title,
		content,
		published,
		priority
	) VALUES (
		$1, $2, $3, $4
	) RETURNING *
`

	var (
		schedule model.Schedule
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&schedule,
		query,
		entity.Title,
		entity.Content,
		entity.Published,
		entity.Priority,
	)

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (repository *Repository) Update(entity *model.Schedule) (*model.Schedule, error) {
	const query = `
    UPDATE schedules
    SET
        title = COALESCE($1, title),
        content = COALESCE($2, content),
        published = COALESCE($3, published),
        priority = COALESCE($4, priority),
        updated_at = $5
    WHERE id = $6
    RETURNING *
`

	var (
		schedule model.Schedule
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&schedule,
		query,
		entity.Title,
		entity.Content,
		entity.Published,
		entity.Priority,
		time.Now(),
		entity.ID,
	)

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (repository *Repository) Find() ([]model.Schedule, error) {
	const query = `
    SELECT
        *
    FROM schedules
	ORDER BY created_at DESC
`

	var (
		schedules []model.Schedule
	)

	err := pgxscan.Select(repository.context, repository.connection.Pool, &schedules, query)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (repository *Repository) First(id uint64) (*model.Schedule, error) {
	const query = `
    SELECT
        *
    FROM schedules
    WHERE id = $1
	ORDER BY created_at DESC
`

	var (
		schedule model.Schedule
	)

	err := pgxscan.Get(repository.context, repository.connection.Pool, &schedule, query, id)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (repository *Repository) Delete(id uint64) error {
	const query = `
	DELETE FROM schedules
	WHERE id = $1
`

	result, err := repository.connection.Exec(repository.context, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return nil
	}

	return nil
}
