package domain

import "context"

type RemoteServer struct {
	ID       int64  `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Name     string `json:"name" db:"name"`
	Address  string `json:"address" db:"address"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type RemoteServerRepository interface {
	Create(ctx context.Context, server *RemoteServer) error
}
