package service

import (
	"context"
	"strconv"
	"time"
)

type service struct {
	serviceRepo    ServiceRepository
	contextTimeout time.Duration
}

func NewService(repository ServiceRepository) ServiceService {
	return &service{
		serviceRepo:    repository,
		contextTimeout: time.Duration(2) * time.Second,
	}
}

func (s *service) CreateService(ctx context.Context, req *ServiceRequest) (*ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	service := &Service{
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
		Status:      req.Status,
	}

	service, err := s.serviceRepo.CreateService(ctx, service)
	if err != nil {
		return nil, err
	}

	return &ServiceResponse{
		ID:          strconv.Itoa(int(service.ID)),
		Name:        service.Name,
		Description: service.Description,
		URL:         service.URL,
		Status:      service.Status,
	}, nil
}

func (s *service) GetServiceByID(ctx context.Context, id string) (*ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	serviceID, err := strconv.Atoi(id)
	if err != nil {
		return &ServiceResponse{}, err
	}

	service, err := s.serviceRepo.GetServiceByID(ctx, int64(serviceID))
	if err != nil {
		return &ServiceResponse{}, err
	}

	return &ServiceResponse{
		ID:          strconv.Itoa(int(service.ID)),
		Name:        service.Name,
		Description: service.Description,
		URL:         service.URL,
		Status:      service.Status,
	}, nil
}
