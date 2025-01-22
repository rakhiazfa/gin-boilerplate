package services

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db             *gorm.DB
	validator      *utils.Validator
	userRepository *repositories.UserRepository
}

func NewAuthService(db *gorm.DB, validator *utils.Validator, userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		db:             db,
		validator:      validator,
		userRepository: userRepository,
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

	if err := s.validateUsername(ctx, req.Username); err != nil {
		return err
	}
	if err := s.validateEmail(ctx, req.Email); err != nil {
		return err
	}

	var user entities.User

	if err := copier.Copy(&user, req); err != nil {
		return err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.userRepository.WithTx(tx).Save(&user)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) validateUsername(ctx context.Context, username string) error {
	user, err := s.userRepository.WithContext(ctx).FindOneByUsernameUnscoped(username)
	if err != nil {
		return err
	}

	if user != nil {
		return utils.NewUniqueFieldError("username", "An account with this username already exists", nil)
	}

	return nil
}

func (s *AuthService) validateEmail(ctx context.Context, email string) error {
	user, err := s.userRepository.WithContext(ctx).FindOneByEmailUnscoped(email)
	if err != nil {
		return err
	}

	if user != nil {
		return utils.NewUniqueFieldError("email", "An account with this email already exists", nil)
	}

	return nil
}
