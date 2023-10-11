package service

import (
	"context"
	"time"
)

type Service struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	URL         string    `json:"url" db:"url"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type CreateServiceResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type ServiceRepository interface {
	CreateService(ctx context.Context, service *Service) (*Service, error)
}

type ServiceService interface {
	CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error)
}
