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
	query := `INSERT INTO remote_servers (user_id, name, address, is_active, last_check_time, next_check_time ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, remoteServer.UserID, remoteServer.Name, remoteServer.Address, remoteServer.IsActive, remoteServer.LastCheckTime, remoteServer.NextCheckTime)
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

func (r *postgresRemoteServerRepo) GetByUserID(ctx context.Context, userID int) ([]domain.RemoteServer, error) {
	query := "SELECT id, name, address, is_active, last_check_time, next_check_time FROM remote_servers WHERE user_id=$1"
	var servers []domain.RemoteServer

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var server domain.RemoteServer
		err := rows.Scan(&server.ID, &server.Name, &server.Address, &server.IsActive, &server.LastCheckTime, &server.NextCheckTime)
		if err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return servers, nil
}

func (r *postgresRemoteServerRepo) GetAll(ctx context.Context) ([]domain.RemoteServer, error) {
	query := "SELECT id, user_id, name, address, is_active, last_check_time, next_check_time FROM remote_servers"
	var servers []domain.RemoteServer

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var server domain.RemoteServer
		err := rows.Scan(&server.ID, &server.UserID, &server.Name, &server.Address, &server.IsActive, &server.LastCheckTime, &server.NextCheckTime)
		if err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return servers, nil
}

func (r *postgresRemoteServerRepo) Update(ctx context.Context, remoteServer *domain.RemoteServer) (err error) {
	query := `UPDATE remote_servers SET name=$1, address=$2, is_active=$3 WHERE id=$4`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, remoteServer.Name, remoteServer.Address, remoteServer.IsActive, remoteServer.ID)
	if err != nil {
		return
	}

	return
}
