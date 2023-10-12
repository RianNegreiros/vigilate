package postgres

import (
	"context"
	"database/sql"

	"github.com/RianNegreiros/vigilate/domain"
)

type postgresUserRepo struct {
	DB *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) domain.UserRepository {
	return &postgresUserRepo{
		DB: db,
	}
}

func (r *postgresUserRepo) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"

	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)
	if err != nil {
		return &domain.User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *postgresUserRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := domain.User{}
	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &domain.User{}, err
	}

	return &u, nil
}
