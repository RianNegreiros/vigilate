package user

import (
	"context"
	"strconv"
	"time"

	"github.com/RianNegreiros/vigilate/util"
)

type service struct {
	userRepo       Repository
	contextTimeout time.Duration
}

func NewService(repository Repository) UserService {
	return &service{
		userRepo:       repository,
		contextTimeout: time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
