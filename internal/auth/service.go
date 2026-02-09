package auth

import (
	"errors"

	"github.com/google/uuid"
)

// Service contains business logic for authentication
type Service struct {
	repo UserRepository
}

type LoginInput struct {
	Email    string
	Password string
}

// NewService creates a new Service instance
func NewService(repo UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

// Register a new user
func (s *Service) Register(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hash,
	}

	return s.repo.Create(user)
}

// Login checks credentials
func (s *Service) Login(input LoginInput) error {
	if input.Email == "" || input.Password == "" {
		return errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	if err := CheckPassword(input.Password, user.Password); err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
