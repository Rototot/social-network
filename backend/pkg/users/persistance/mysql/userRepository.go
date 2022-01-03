package mysql

import (
	"context"
	"database/sql"
	"errors"
	"social-network/pkg/users"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindById(ctx context.Context, id users.UserID) (*users.User, error) {
	var user users.User
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM users WHEN id = :id",
		sql.Named("id", id),
	).Scan(
		&user.ID,
		&user.Email,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Gender,
		&user.City,
		&user.Interests,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmailAndPassword(
	ctx context.Context,
	email string,
	password users.HashedPassword,
) (*users.User, error) {
	var user users.User
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM users WHEN email = :email, password = :password",
		sql.Named("email", email),
		sql.Named("password", password),
	).Scan(
		&user.ID,
		&user.Email,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Gender,
		&user.City,
		&user.Interests,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
