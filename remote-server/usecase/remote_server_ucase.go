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

	remoteServer := &domain.RemoteServer{
		UserID:   req.UserID,
		Name:     req.Name,
		Address:  req.Address,
		IsActive: req.IsActive,
	}

	err = s.remoteServerRepo.Create(ctx, remoteServer)

	return
}
