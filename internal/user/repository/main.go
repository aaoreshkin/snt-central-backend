package repository

import (
	"context"

	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/user/entity"
)

type UserRepository struct {
	context    context.Context
	connection *database.Connection
}

func NewUserRepository(context context.Context, connection *database.Connection) *UserRepository {
	return &UserRepository{context, connection}
}

func (repository *UserRepository) Create(entity *entity.User) (*entity.User, error) {

	query := `
		INSERT INTO users (
			access_token,
			email,
			password,
			permission,
			fullname,
			phone
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	repository.connection.Pool.Exec(repository.context, query, entity.AccessToken, entity.Email, entity.Password, entity.Permission, entity.Fullname, entity.Phone)

	return entity, nil
}

func (repository *UserRepository) Find() ([]entity.User, error) {
	entities := []entity.User{}

	query := `
		SELECT *
		FROM users
	`

	rows, err := repository.connection.Pool.Query(repository.context, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.AccessToken, &user.Email, &user.Password, &user.Permission, &user.Fullname, &user.Phone, &user.UpdatedAt); err != nil {
			return nil, err
		}
		entities = append(entities, user)
	}

	return entities, nil
}

func (repository *UserRepository) First(email string) (*entity.User, error) {
	entity := &entity.User{}

	query := `
		SELECT *
		FROM users
		WHERE email = $1;
	`

	if err := repository.connection.QueryRow(repository.context, query, email).Scan(&entity.ID, &entity.AccessToken, &entity.Email, &entity.Password, &entity.Permission, &entity.Fullname, &entity.Phone, &entity.UpdatedAt); err != nil {
		// Return an error if the product is not found.
		return nil, err
	}

	return entity, nil
}

func (repository *UserRepository) Update(entity *entity.User, id uint64) (*entity.User, error) {

	query := `
		UPDATE users
		SET
			access_token = $1,
			email = $2,
			password = $3,
			permission = $4,
			fullname = $5,
			phone = $6
		WHERE id = $7
	`

	repository.connection.Pool.Exec(repository.context, query, entity.AccessToken, entity.Email, entity.Password, entity.Permission, entity.Fullname, entity.Phone, id)

	return entity, nil
}

func (repository *UserRepository) Delete(id uint64) error {

	query := `
		DELETE FROM users
		WHERE id = $1
	`

	repository.connection.Pool.Exec(repository.context, query, id)

	return nil
}
