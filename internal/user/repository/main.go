package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/user/model"
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

func (repository *Repository) Create(entity *model.User) (*model.User, error) {
	const query = `
	INSERT INTO users (
		access_token,
		description,
		email,
		fullname,
		password,
		permission_id,
		phone,
		position,
		refresh_token
	) VALUES (
	 	$1, $2, $3, $4, $5, $6, $7, $8, $9
	) RETURNING
		id,
		access_token,
		description,
		email,
		fullname,
		permission_id,
		phone,
		position,
		refresh_token,
		created_at,
		updated_at
`

	var (
		user model.User
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&user,
		query,
		entity.AccessToken,
		entity.Description,
		entity.Email,
		entity.Fullname,
		entity.Password,
		entity.PermissionID,
		entity.Phone,
		entity.Position,
		entity.RefreshToken,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) Find() ([]model.User, error) {
	const query = `
    SELECT
		id,
		access_token,
		description,
		email,
		fullname,
		permission_id,
		phone,
		position,
		refresh_token,
		created_at,
		updated_at
    FROM users
	ORDER BY created_at DESC
`

	var (
		users []model.User
	)

	err := pgxscan.Select(repository.context, repository.connection.Pool, &users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *Repository) First(id uint64) (*model.User, error) {
	const query = `
    SELECT
		id,
		access_token,
		description,
		email,
		fullname,
		permission_id,
		phone,
		position,
		refresh_token,
		created_at,
		updated_at
    FROM users
    WHERE id = $1
	ORDER BY created_at DESC
`

	var (
		user model.User
	)

	err := pgxscan.Get(repository.context, repository.connection.Pool, &user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) Any(key, value string) (*model.User, error) {
	query := fmt.Sprintf(`
        SELECT
            *
        FROM users
        WHERE %s = $1;
    `, key)

	var (
		user model.User
	)

	err := pgxscan.Get(repository.context, repository.connection.Pool, &user, query, value)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) Update(entity *model.User) (*model.User, error) {
	const query = `
    UPDATE users
    SET
        access_token = COALESCE($1, access_token),
        description = COALESCE($2, description),
        email = COALESCE($3, email),
        fullname = COALESCE($4, fullname),
        password = COALESCE($5, password),
        permission_id = COALESCE($6, permission_id),
        phone = COALESCE($7, phone),
        position = COALESCE($8, position),
        refresh_token = COALESCE($9, refresh_token),
        updated_at = $10
    WHERE id = $11
	RETURNING
		id,
		access_token,
		description,
		email,
		fullname,
		permission_id,
		phone,
		position,
		refresh_token,
		created_at,
		updated_at
`

	var (
		user model.User
	)

	err := pgxscan.Get(
		repository.context,
		repository.connection,
		&user,
		query,
		entity.AccessToken,
		entity.Description,
		entity.Email,
		entity.Fullname,
		entity.Password,
		entity.PermissionID,
		entity.Phone,
		entity.Position,
		entity.RefreshToken,
		time.Now(),
		entity.ID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) Delete(id uint64) error {
	const query = `
	DELETE FROM users
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
