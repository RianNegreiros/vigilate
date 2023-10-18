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
	query := `INSERT INTO remote_servers (user_id, name, address, is_active) VALUES ($1, $2, $3, $4) RETURNING id`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, remoteServer.UserID, remoteServer.Name, remoteServer.Address, remoteServer.IsActive)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	remoteServer.ID = lastID
	return
}
