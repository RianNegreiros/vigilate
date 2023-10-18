package domain

import "context"

type RemoteServer struct {
	ID       int64  `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Name     string `json:"name" db:"name"`
	Address  string `json:"address" db:"address"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type CreateRemoteServer struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	IsActive bool   `json:"is_active"`
}

type RemoteServerRepository interface {
	Create(ctx context.Context, server *RemoteServer) error
}

type RemoteServerUsecase interface {
	Create(ctx context.Context, server *CreateRemoteServer) error
}
