package repository

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/events/model"
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

func (repository *Repository) Create(entity *model.Event) (*model.Event, error) {
	const query = `
	INSERT INTO events (
		title,
		published,
		created_at
	) VALUES (
		$1, $2, $3
	) RETURNING *
`

	var (
		event model.Event
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&event,
		query,
		entity.Title,
		entity.Published,
		entity.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Если есть файлы, добавляем их
	if len(entity.Attachments) > 0 {
		const query = `
		INSERT INTO events_attachments (event_id, file_id)
		VALUES ($1, $2)
`
		for _, file := range entity.Attachments {
			_, err = repository.connection.Pool.Exec(
				repository.context,
				query,
				event.ID,
				file.ID,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	return &event, nil
}

func (repository *Repository) Update(entity *model.Event) (*model.Event, error) {
	const query = `
    UPDATE events
    SET
        title = COALESCE($1, title),
        published = COALESCE($2, published),
        created_at = COALESCE($3, created_at),
        updated_at = $4
    WHERE id = $5
    RETURNING *
`

	var (
		event model.Event
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&event,
		query,
		entity.Title,
		entity.Published,
		entity.CreatedAt,
		time.Now(),
		entity.ID,
	)

	if err != nil {
		return nil, err
	}

	// Сначала удаляем все существующие связи
	deleteQuery := `DELETE FROM events_attachments WHERE event_id = $1`
	_, err = repository.connection.Exec(repository.context, deleteQuery, entity.ID)
	if err != nil {
		return nil, err
	}

	// Добавляем новые связи
	if len(entity.Attachments) > 0 {
		insertQuery := `
            INSERT INTO events_attachments (event_id, file_id, created_at, updated_at)
            VALUES ($1, $2, $3, $3)
        `
		now := time.Now()
		for _, file := range entity.Attachments {
			_, err = repository.connection.Exec(
				repository.context,
				insertQuery,
				entity.ID,
				file.ID,
				now,
			)
			if err != nil {
				return nil, err
			}
		}
	}

	return &event, nil
}

func (repository *Repository) Find() ([]model.Event, error) {
	const query = `
    SELECT
        e.*,
        COALESCE(json_agg(
            json_build_object(
                'id', f.id,
                'name', f.name,
                'hex', f.hex,
                'size', f.size,
                'extension', f.extension,
                'created_at', f.created_at,
                'updated_at', f.updated_at
            ) ORDER BY f.created_at
        ) FILTER (WHERE f.id IS NOT NULL), '[]') as attachments
    FROM events e
    LEFT JOIN events_attachments ef ON e.id = ef.event_id
    LEFT JOIN attachments f ON ef.file_id = f.id
	GROUP BY e.id, e.title, e.published, e.created_at, e.updated_at
	ORDER BY e.created_at DESC
`

	var (
		events []model.Event
	)

	err := pgxscan.Select(repository.context, repository.connection.Pool, &events, query)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (repository *Repository) First(id uint64) (*model.Event, error) {
	const query = `
    SELECT
        e.*,
        COALESCE(json_agg(
            json_build_object(
                'id', f.id,
                'name', f.name,
                'hex', f.hex,
                'size', f.size,
                'extension', f.extension,
                'created_at', f.created_at,
                'updated_at', f.updated_at
            ) ORDER BY f.created_at
        ) FILTER (WHERE f.id IS NOT NULL), '[]') as attachments
    FROM events e
    LEFT JOIN events_attachments ef ON e.id = ef.event_id
    LEFT JOIN attachments f ON ef.file_id = f.id
    WHERE e.id = $1
    GROUP BY e.id, e.title, e.published, e.created_at, e.updated_at
`

	var (
		event model.Event
	)

	err := pgxscan.Get(repository.context, repository.connection.Pool, &event, query, id)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (repository *Repository) Delete(id uint64) error {
	const query = `
	DELETE FROM events
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
