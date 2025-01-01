package services

import (
	"context"

	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	validator *utils.Validator
}

func NewAuthService(db *gorm.DB, validator *utils.Validator) *AuthService {
	return &AuthService{
		db:        db,
		validator: validator,
	}
}

func (s *AuthService) SignIn(ctx context.Context, req *dtos.SignInReq) (string, error) {
	if err := s.validator.Validate(req); err != nil {
		return "", err
	}

	return "", nil
}

func (s *AuthService) SignUp(ctx context.Context, req *dtos.SignUpReq) error {
	if err := s.validator.Validate(req); err != nil {
		return err
	}

	return nil
}
