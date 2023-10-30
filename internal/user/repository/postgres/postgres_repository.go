package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
	query := "SELECT id, email, username, notification_preferences->>'email_enabled', created_at, updated_at FROM users WHERE id = $1"
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Email, &u.Username, &u.NotificationPreferences.EmailEnabled, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (r *postgresUserRepo) updateNotificationPreferences(ctx context.Context, userID int, preferences interface{}, notificationType string) error {
	var notificationPreferencesJSON []byte
	query := "SELECT notification_preferences FROM users WHERE id = $1"
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&notificationPreferencesJSON)
	if err != nil {
		log.Printf("Error fetching notification preferences for %s: %v\n", notificationType, err)
		return err
	}

	var notificationPreferences domain.NotificationPreferences
	err = json.Unmarshal(notificationPreferencesJSON, &notificationPreferences)
	if err != nil {
		log.Printf("Error unmarshaling JSON for %s: %v\n", notificationType, err)
		return err
	}

	switch notificationType {
	case "email":
		notificationPreferences.EmailEnabled = preferences.(bool)
	default:
		log.Printf("Unknown notification type: %s\n", notificationType)
		return errors.New("unknown notification type")
	}

	updatedNotificationPreferencesJSON, err := json.Marshal(notificationPreferences)
	if err != nil {
		log.Printf("Error marshaling JSON for %s: %v\n", notificationType, err)
		return err
	}

	updateQuery := "UPDATE users SET notification_preferences = $1 WHERE id = $2"
	_, err = r.DB.ExecContext(ctx, updateQuery, updatedNotificationPreferencesJSON, userID)
	if err != nil {
		log.Printf("Error updating notification preferences for %s: %v\n", notificationType, err)
		return err
	}

	return nil
}

func (r *postgresUserRepo) UpdateEmailNotificationPreferences(ctx context.Context, userID int, emailEnabled bool) error {
	return r.updateNotificationPreferences(ctx, userID, emailEnabled, "email")
}
