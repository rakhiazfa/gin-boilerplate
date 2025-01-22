package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{
		db: tx,
	}
}

func (r *UserRepository) WithContext(ctx context.Context) *UserRepository {
	return &UserRepository{
		db: r.db.WithContext(ctx),
	}
}

func (r *UserRepository) Save(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) FindOneByUsernameUnscoped(username string, excludedIds ...uuid.UUIDs) (*entities.User, error) {
	var user entities.User

	q := r.db.Unscoped().Where("username = ?", username)

	if len(excludedIds) > 0 {
		q = q.Where("id NOT IN ?", excludedIds)
	}

	if err := q.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindOneByEmailUnscoped(email string, excludedIds ...uuid.UUIDs) (*entities.User, error) {
	var user entities.User

	q := r.db.Unscoped().Where("email = ?", email)

	if len(excludedIds) > 0 {
		q = q.Where("id NOT IN ?", excludedIds)
	}

	if err := q.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
