package usecase

import (
	"context"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	// Check if the provided email is valid
	if !isValidEmail(req.Email) {
		return nil, domain.ErrInvalidEmail
	}

	// Check if the email is unique
	userExists, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userExists.Email != "" {
		return nil, domain.ErrDuplicateEmail
	}

	// Ensure that the password meets your strength requirements (e.g., length, special characters)
	if !isStrongPassword(req.Password) {
		return nil, domain.ErrWeakPassword
	}

	// Verify that the "Password" and "ConfirmPassword" fields match
	if req.Password != req.ConfirmPassword {
		return nil, domain.ErrPasswordMismatch
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
	Email    string `json:"email"`
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
		Email:    u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return &domain.LoginUserResponse{}, err
	}

	return &domain.LoginUserResponse{AccessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID)), Email: u.Email}, nil
}

func (s *userUsecase) UpdateEmailNotificationPreferences(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("Error getting user: ", err)
		return err
	}

	emailEnabled := !user.NotificationPreferences.EmailEnabled

	err = s.userRepo.UpdateEmailNotificationPreferences(ctx, userID, emailEnabled)
	if err != nil {
		log.Println("Error updating notification preferences: ", err)
		return err
	}

	return nil
}

func (s *userUsecase) GetByID(ctx context.Context, userID int) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("Error getting user: ", err)
		return nil, err
	}

	return user, nil
}

func isValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func isStrongPassword(password string) bool {
	if len(password) < 8 {
		return false // Minimum length requirement not met
	}

	// Check for at least one uppercase letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	// Check for at least one lowercase letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	// Check for at least one special character (you can customize this as needed)
	specialCharacters := "!@#$%^&*()-_=+[]{}|;:'\",<.>/?"
	for _, char := range password {
		if strings.ContainsRune(specialCharacters, char) {
			return true
		}
	}

	return false
}
