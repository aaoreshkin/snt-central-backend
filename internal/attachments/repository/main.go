package repository

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/oreshkindev/snt-central-backend/internal/attachments/model"
	"github.com/oreshkindev/snt-central-backend/internal/database"
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

func (repository *Repository) Create(entity *model.Attachment) (*model.Attachment, error) {
	const query = `
	INSERT INTO attachments (
		name,
		hex,
		size,
		extension
	) VALUES (
		$1, $2, $3, $4
	) RETURNING *
`

	var (
		attachment model.Attachment
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&attachment,
		query,
		entity.Name,
		entity.Hex,
		entity.Size,
		entity.Extension,
	)

	if err != nil {
		return nil, err
	}

	return &attachment, nil
}

func (repository *Repository) Find() ([]*model.Attachment, error) {
	const query = `
        SELECT
            *
        FROM attachments
        ORDER BY created_at DESC
`

	var (
		attachments []*model.Attachment
	)

	err := pgxscan.Select(
		repository.context,
		repository.connection,
		&attachments,
		query,
	)
	if err != nil {
		return nil, err
	}

	return attachments, nil
}

func (repository *Repository) Delete(id uint64) error {
	const query = `
	DELETE FROM attachments
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
