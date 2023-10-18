package usecase

import (
	"context"
	"time"

	"github.com/RianNegreiros/vigilate/domain"
)

type remoteServerUsecase struct {
	remoteServerRepo domain.RemoteServerRepository
	contextTimeout   time.Duration
}

func NewRemoteServerUsecase(u domain.RemoteServerRepository, timeout time.Duration) domain.RemoteServerUsecase {
	return &remoteServerUsecase{
		remoteServerRepo: u,
		contextTimeout:   timeout,
	}
}

func (s *remoteServerUsecase) Create(ctx context.Context, req *domain.CreateRemoteServer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	exists, err := s.remoteServerRepo.Exists(ctx, req.Address)
	if err != nil {
		return
	}

	if exists {
		return domain.ErrDuplicateAddress
	}

	remoteServer := &domain.RemoteServer{
		UserID:   req.UserID,
		Name:     req.Name,
		Address:  req.Address,
		IsActive: req.IsActive,
	}

	err = s.remoteServerRepo.Create(ctx, remoteServer)

	return
}

func (s *remoteServerUsecase) GetByUserID(ctx context.Context, userID int) (servers []domain.RemoteServer, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	servers, err = s.remoteServerRepo.GetByUserID(ctx, userID)

	return servers, err
}
