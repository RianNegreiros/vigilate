package domain

import (
	"context"
	"time"
)

type User struct {
	ID                      int64                   `json:"id" db:"id"`
	Username                string                  `json:"username" db:"username"`
	Email                   string                  `json:"email" db:"email"`
	Password                []byte                  `json:"password" db:"password"`
	CreatedAt               time.Time               `json:"created_at" db:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at" db:"updated_at"`
	NotificationPreferences NotificationPreferences `json:"notification_preferences" db:"notification_preferences"`
}

type NotificationPreferences struct {
	EmailEnabled bool `json:"email_enabled" db:"email_enabled"`
}

type CreateUserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	UpdateEmailNotificationPreferences(ctx context.Context, userID int, emailEnabled bool) error
}

type UserUsecase interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	Login(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
	UpdateEmailNotificationPreferences(ctx context.Context, userID int) error
	GetByID(ctx context.Context, userID int) (*User, error)
}
