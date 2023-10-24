package domain

import (
	"context"
	"time"
)

type User struct {
	ID          int64             `json:"id" db:"id"`
	Username    string            `json:"username" db:"username"`
	Email       string            `json:"email" db:"email"`
	Password    []byte            `json:"password" db:"password"`
	Preferences map[string]string `json:"preferences" db:"preferences"`
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
}

type UserUsecase interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	Login(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
}
