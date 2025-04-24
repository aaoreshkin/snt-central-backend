package repository

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/posts/model"
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

func (repository *Repository) Create(entity *model.Post) (*model.Post, error) {
	const query = `
	INSERT INTO posts (
		title,
		content,
		published,
		ads,
		created_at
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING *
`

	var (
		post model.Post
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&post,
		query,
		entity.Title,
		entity.Content,
		entity.Published,
		entity.Ads,
		entity.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (repository *Repository) Update(entity *model.Post) (*model.Post, error) {
	const query = `
    UPDATE posts
    SET
        title = COALESCE($1, title),
        content = COALESCE($2, content),
        published = COALESCE($3, published),
        ads = COALESCE($4, ads),
        created_at = COALESCE($5, created_at),
        updated_at = $6
    WHERE id = $7
    RETURNING *
`

	var (
		post model.Post
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&post,
		query,
		entity.Title,
		entity.Content,
		entity.Published,
		entity.Ads,
		entity.CreatedAt,
		time.Now(),
		entity.ID,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (repository *Repository) Find() ([]model.Post, error) {
	const query = `
    SELECT
        *
    FROM posts
	ORDER BY created_at DESC
`

	var (
		posts []model.Post
	)

	err := pgxscan.Select(repository.context, repository.connection.Pool, &posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *Repository) First(id uint64) (*model.Post, error) {
	const query = `
    SELECT
        *
    FROM posts
    WHERE id = $1
	ORDER BY created_at DESC
`

	var (
		post model.Post
	)

	err := pgxscan.Get(repository.context, repository.connection.Pool, &post, query, id)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (repository *Repository) Delete(id uint64) error {
	const query = `
	DELETE FROM posts
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
