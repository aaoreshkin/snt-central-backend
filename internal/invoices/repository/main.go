package repository

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/invoices/model"
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

func (repository *Repository) Create(entity *model.Invoice) (*model.Invoice, error) {
	const query = `
	INSERT INTO invoices (
		title,
		content,
		published,
		priority,
		kwt_price
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING *
`

	var (
		schedule model.Invoice
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
		entity.KwtPrice,
	)

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (repository *Repository) Update(entity *model.Invoice) (*model.Invoice, error) {
	const query = `
    UPDATE invoices
    SET
        title = COALESCE($1, title),
        content = COALESCE($2, content),
        published = COALESCE($3, published),
        priority = COALESCE($4, priority),
		kwt_price = COALESCE($5, kwt_price),
        updated_at = $6
    WHERE id = $7
    RETURNING *
`

	var (
		schedule model.Invoice
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
		entity.KwtPrice,
		time.Now(),
		entity.ID,
	)

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (repository *Repository) Find() ([]model.Invoice, error) {
	const query = `
    SELECT
        *
    FROM invoices
	ORDER BY created_at DESC
`

	var (
		invoices []model.Invoice
	)

	err := pgxscan.Select(repository.context, repository.connection.Pool, &invoices, query)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (repository *Repository) First(id uint64) (*model.Invoice, error) {
	const query = `
    SELECT
        *
    FROM invoices
    WHERE id = $1
	ORDER BY created_at DESC
`

	var (
		schedule model.Invoice
	)

	err := pgxscan.Get(repository.context, repository.connection.Pool, &schedule, query, id)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (repository *Repository) Delete(id uint64) error {
	const query = `
	DELETE FROM invoices
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
