package user

import (
	"context"
	"strconv"
	"time"

	"github.com/RianNegreiros/vigilate/util"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
	domain    = "localhost"
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

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	u, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserResponse{}, err
	}

	return &LoginUserResponse{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}
