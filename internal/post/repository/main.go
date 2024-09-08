package repository

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/post/entity"
)

type PostRepository struct {
	context    context.Context
	connection *database.Connection
}

func NewPostRepository(context context.Context, connection *database.Connection) *PostRepository {
	return &PostRepository{context, connection}
}

func (repository *PostRepository) Create(entity *entity.Post) (*entity.Post, error) {

	query := `
		INSERT INTO posts (
			title,
			description,
			promo,
			published
		)
		VALUES ($1, $2, $3, $4)
	`

	repository.connection.Pool.Exec(repository.context, query, entity.Title, entity.Description, entity.Promo, entity.Published)

	return entity, nil
}

func (repository *PostRepository) Find() ([]entity.Post, error) {

	query := `
		SELECT *
		FROM posts
	`

	entities := []entity.Post{}

	rows, err := repository.connection.Pool.Query(repository.context, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.Post
		if err := rows.Scan(&entity.ID, &entity.Promo, &entity.Published, &entity.Title, &entity.Description, &entity.CreatedAt); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (repository *PostRepository) First(id string) (*entity.Post, error) {

	query := `
		SELECT *
		FROM posts
		WHERE id = $1;
	`

	entity := &entity.Post{}

	if err := repository.connection.QueryRow(repository.context, query, id).Scan(&entity.ID, &entity.Promo, &entity.Published, &entity.Title, &entity.Description, &entity.CreatedAt); err != nil {
		// Return an error if the product is not found.
		return nil, err
	}

	return entity, nil
}

func (repository *PostRepository) Update(entity *entity.Post, id string) (*entity.Post, error) {

	query := `
		UPDATE posts
		SET
			title = $1,
			description = $2,
			promo = $3,
			published = $4
		WHERE id = $5
	`

	repository.connection.Pool.Exec(repository.context, query, entity.Title, entity.Description, entity.Promo, entity.Published, id)

	return entity, nil
}

func (repository *PostRepository) Delete(id string) error {

	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	repository.connection.Pool.Exec(repository.context, query, id)

	return nil
}
