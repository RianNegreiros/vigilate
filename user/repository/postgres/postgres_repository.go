package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

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
		log.Println("Error executing statement: ", err)
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
		log.Println("Error executing statement: ", err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (r *postgresUserRepo) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	u := domain.User{}
	query := "SELECT id, email, username, password FROM users WHERE id = $1"
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (r *postgresUserRepo) AllPreferences(ctx context.Context) ([]domain.Preference, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "SELECT id, name, preference FROM preferences"

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return nil, err
	}
	defer rows.Close()

	var preferences []domain.Preference

	for rows.Next() {
		s := &domain.Preference{}
		err := rows.Scan(&s.ID, &s.Name, &s.Preference)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return nil, err
		}
		preferences = append(preferences, *s)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error scanning row: ", err)
		return nil, err
	}

	return preferences, nil
}

func (r *postgresUserRepo) SetSystemPref(ctx context.Context, name, value string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	stmt := `delete from preferences where name = $1`
	_, _ = r.DB.ExecContext(ctx, stmt, name)

	query := `INSERT INTO preferences(name, preference) VALUES ($1, $2)`

	_, err := r.DB.ExecContext(ctx, query, name, value)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return err
	}

	return nil
}

func (m *postgresUserRepo) InsertOrUpdateSitePreferences(ctx context.Context, pm map[string]string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	for k, v := range pm {
		query := `delete from preferences where name = $1`

		_, err := m.DB.ExecContext(ctx, query, k)
		if err != nil {
			log.Println("Error executing statement: ", err)
			return err
		}

		query = `insert into preferences (name, preference) values ($1, $2)`

		_, err = m.DB.ExecContext(ctx, query, k, v)
		if err != nil {
			log.Println("Error executing statement: ", err)
			return err
		}
	}

	return nil
}
