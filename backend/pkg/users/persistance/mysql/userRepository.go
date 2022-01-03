package mysql

import (
	"context"
	"database/sql"
	"social-network/pkg/users"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) FindByEmailAndPassword(ctx context.Context, email string, password users.HashedPassword) (*users.User, error) {
	var user users.User
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, email FROM users WHEN email = :email, password = :password",
		sql.Named("email", email),
		sql.Named("password", password),
	).Scan(&user.ID, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
