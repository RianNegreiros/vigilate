package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/RianNegreiros/vigilate/internal/domain"
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
	query := `INSERT INTO remote_servers (user_id, name, address, is_active, last_check_time, next_check_time, notified) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, remoteServer.UserID, remoteServer.Name, remoteServer.Address, remoteServer.IsActive, remoteServer.LastCheckTime, remoteServer.NextCheckTime, remoteServer.Notified)
	if err != nil {
		log.Println("Error executing statement: ", err)
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
		log.Println("Error executing statement: ", err)
		return false, err
	}

	return exists, nil
}

func (r *postgresRemoteServerRepo) GetByUserID(ctx context.Context, userID int) ([]domain.RemoteServer, error) {
	stmt, err := r.DB.PrepareContext(ctx, "SELECT id, name, user_id, address, is_active, last_check_time, next_check_time FROM remote_servers WHERE user_id=$1")
	if err != nil {
		log.Println("Error preparing statement: ", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return nil, err
	}
	defer rows.Close()

	var servers []domain.RemoteServer
	for rows.Next() {
		var server domain.RemoteServer
		err := rows.Scan(&server.ID, &server.Name, &server.UserID, &server.Address, &server.IsActive, &server.LastCheckTime, &server.NextCheckTime)
		if err != nil {
			log.Println("Error scanning rows: ", err)
			return nil, err
		}
		servers = append(servers, server)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows: ", err)
		return nil, err
	}

	return servers, nil
}

func (r *postgresRemoteServerRepo) GetAll(ctx context.Context) ([]domain.RemoteServer, error) {
	query := "SELECT id, user_id, name, address, is_active, last_check_time, next_check_time FROM remote_servers"
	var servers []domain.RemoteServer

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var server domain.RemoteServer
		err := rows.Scan(&server.ID, &server.UserID, &server.Name, &server.Address, &server.IsActive, &server.LastCheckTime, &server.NextCheckTime)
		if err != nil {
			log.Println("Error scanning rows: ", err)
			return nil, err
		}
		servers = append(servers, server)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows: ", err)
		return nil, err
	}

	return servers, nil
}

func (r *postgresRemoteServerRepo) Update(ctx context.Context, remoteServer *domain.RemoteServer) (err error) {
	query := `UPDATE remote_servers SET name=$1, address=$2, is_active=$3, last_check_time=$4, next_check_time=$5, notified=$6 WHERE id=$7`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, remoteServer.Name, remoteServer.Address, remoteServer.IsActive, remoteServer.LastCheckTime, remoteServer.NextCheckTime, remoteServer.Notified, remoteServer.ID)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return
	}

	return
}

func (r *postgresRemoteServerRepo) GetByID(ctx context.Context, id int) (domain.RemoteServer, error) {
	query := "SELECT id, user_id, name, address, is_active, last_check_time, next_check_time FROM remote_servers WHERE id=$1"
	var server domain.RemoteServer

	err := r.DB.QueryRowContext(ctx, query, id).Scan(&server.ID, &server.UserID, &server.Name, &server.Address, &server.IsActive, &server.LastCheckTime, &server.NextCheckTime)
	if err != nil {
		log.Println("Error executing statement: ", err)
		return domain.RemoteServer{}, err
	}

	return server, nil
}

func (r *postgresRemoteServerRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM remote_servers WHERE id=$1"
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Error preparing statement: ", err)
	}

	_, err = stmt.ExecContext(ctx, id)

	return err
}
