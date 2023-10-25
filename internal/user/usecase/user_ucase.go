package usecase

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RianNegreiros/vigilate/internal/domain"
	"github.com/RianNegreiros/vigilate/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (s *userUsecase) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	userExists, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userExists.Email != "" {
		return nil, domain.ErrDuplicateEmail
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Println("Error hashing password: ", err)
		return nil, err
	}

	user := &domain.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error creating user: ", err)
		return nil, err
	}

	return &domain.CreateUserResponse{
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

func (s *userUsecase) Login(c context.Context, req *domain.LoginUserRequest) (*domain.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	u, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &domain.LoginUserResponse{}, domain.ErrNoRecord
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return &domain.LoginUserResponse{}, domain.ErrInvalidCredentials
	} else if err != nil {
		return &domain.LoginUserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return &domain.LoginUserResponse{}, err
	}

	return &domain.LoginUserResponse{AccessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

func (s *userUsecase) UpdateNotificationPreferences(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("Error getting user: ", err)
		return err
	}

	emailEnabled := !user.NotificationPreferences.EmailEnabled

	err = s.userRepo.UpdateNotificationPreferences(ctx, userID, emailEnabled)
	if err != nil {
		log.Println("Error updating notification preferences: ", err)
		return err
	}

	return nil
}