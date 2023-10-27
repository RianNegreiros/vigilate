package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/RianNegreiros/vigilate/internal/domain"
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
	query := "SELECT id, email, username, notification_preferences->>'email_enabled, created_at updated_at FROM users WHERE id = $1"
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Email, &u.Username, &u.NotificationPreferences.EmailEnabled, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (r *postgresUserRepo) UpdateNotificationPreferences(ctx context.Context, userID int, emailEnabled bool) error {
	var notificationPreferencesJSON []byte
	query := "SELECT notification_preferences FROM users WHERE id = $1"
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&notificationPreferencesJSON)
	if err != nil {
		log.Println("Error fetching notification preferences: ", err)
		return err
	}

	var notificationPreferences domain.NotificationPreferences
	err = json.Unmarshal(notificationPreferencesJSON, &notificationPreferences)
	if err != nil {
		log.Println("Error unmarshaling JSON: ", err)
		return err
	}

	notificationPreferences.EmailEnabled = emailEnabled

	updatedNotificationPreferencesJSON, err := json.Marshal(notificationPreferences)
	if err != nil {
		log.Println("Error marshaling JSON: ", err)
		return err
	}

	updateQuery := "UPDATE users SET notification_preferences = $1 WHERE id = $2"
	_, err = r.DB.ExecContext(ctx, updateQuery, updatedNotificationPreferencesJSON, userID)
	if err != nil {
		log.Println("Error updating notification preferences: ", err)
		return err
	}

	return nil
}
