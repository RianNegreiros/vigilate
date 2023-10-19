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
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type RemoteServerRepository interface {
	Create(ctx context.Context, server *RemoteServer) error
	Exists(ctx context.Context, address string) (bool, error)
	GetByUserID(ctx context.Context, userID int) ([]RemoteServer, error)
}

type RemoteServerUsecase interface {
	Create(ctx context.Context, server *CreateRemoteServer) error
	GetByUserID(ctx context.Context, userID int) ([]RemoteServer, error)
}
