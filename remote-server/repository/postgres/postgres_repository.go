package postgres

import (
	"context"
	"database/sql"

	"github.com/RianNegreiros/vigilate/domain"
)

type postgresRemoteServerRepo struct {
	DB *sql.DB
}

func NewPostgresRemoteServerRepo(db *sql.DB) domain.RemoteServerRepository {
	return &postgresRemoteServerRepo{
		DB: db,
	}
}

func (r *postgresRemoteServerRepo) Create(ctx context.Context, remoteServer *domain.RemoteServer) (err error) {
	var lastInsertId int
	query := `INSERT INTO remote_servers (user_id, name, address, is_active) VALUES ($1, $2, $3, $4) RETURNING id`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, remoteServer.UserID, remoteServer.Name, remoteServer.Address, remoteServer.IsActive)
	if err != nil {
		return
	}

	remoteServer.ID = int64(lastInsertId)
	return
}

func (r *postgresRemoteServerRepo) Exists(ctx context.Context, address string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM remote_servers WHERE address=$1)`
	var exists bool
	err := r.DB.QueryRowContext(ctx, query, address).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
