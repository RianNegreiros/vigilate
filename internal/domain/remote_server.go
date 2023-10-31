package domain

import (
	"context"
	"time"
)

type RemoteServer struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int       `json:"user_id" db:"user_id"`
	Name          string    `json:"name" db:"name"`
	Address       string    `json:"address" db:"address"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	LastCheckTime time.Time `json:"last_check_time" db:"last_check_time"`
	NextCheckTime time.Time `json:"next_check_time" db:"next_check_time"`
	Notified      bool      `json:"notified" db:"notified"`
}

type CreateRemoteServer struct {
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateRemoteServer struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type RemoteServerRepository interface {
	Create(ctx context.Context, server *RemoteServer) error
	Exists(ctx context.Context, address string) (bool, error)
	GetByID(ctx context.Context, id int) (RemoteServer, error)
	GetByUserID(ctx context.Context, userID int) ([]RemoteServer, error)
	GetAll(ctx context.Context) ([]RemoteServer, error)
	Update(ctx context.Context, server *RemoteServer) error
	UpdateNameAddress(ctx context.Context, server *UpdateRemoteServer) (*RemoteServer, error)
	Delete(ctx context.Context, id int) error
}

type RemoteServerUsecase interface {
	Create(ctx context.Context, server *CreateRemoteServer) error
	GetByUserID(ctx context.Context, userID int) ([]RemoteServer, error)
	GetByID(ctx context.Context, id int) (RemoteServer, error)
	StartMonitoring(serverID int) error
	UpdateNameAddress(ctx context.Context, server *UpdateRemoteServer) (*RemoteServer, error)
	Delete(ctx context.Context, id int) error
}
