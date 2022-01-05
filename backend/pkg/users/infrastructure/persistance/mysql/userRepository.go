package mysql

import (
	"context"
	"database/sql"
	"errors"
	"social-network/pkg/users"
	"social-network/pkg/users/services"
)

type UserRepository struct {
	services.UserRepositoryInterface
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(ctx context.Context, id users.UserID) (*users.User, error) {
	var user users.User
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM users WHERE id = ?",
		id,
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
		`
	   SELECT id, email,  password, first_name, last_name, age, gender, city, interests
	   FROM users
	   WHERE email = ? AND password = ?`,
		email,
		password,
	).Scan(
		&user.ID,
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
