package postgres

import (
	"context"
	"database/sql"

	"github.com/RianNegreiros/vigilate/domain"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) domain.ServiceRepository {
	return &repository{db: db}
}

func (r *repository) CreateService(ctx context.Context, service *domain.Service) (*domain.Service, error) {
	var lastInsertId int
	query := "INSERT INTO services(name, description, url, status) VALUES ($1, $2, $3, $4) returning id"

	err := r.db.QueryRowContext(ctx, query, service.Name, service.Description, service.URL, service.Status).Scan(&lastInsertId)
	if err != nil {
		return &domain.Service{}, err
	}

	service.ID = int64(lastInsertId)
	return service, nil
}

func (r *repository) GetServiceByID(ctx context.Context, id int64) (*domain.Service, error) {
	s := domain.Service{}
	query := "SELECT id, name, description, url, status FROM services WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&s.ID, &s.Name, &s.Description, &s.URL, &s.Status)
	if err != nil {
		return &domain.Service{}, err
	}

	return &s, nil
}
