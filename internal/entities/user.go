package entities

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ApplicationEntity
	ProfilePicture uuid.UUID `gorm:"type:uuid;default:null"`
	Name           string    `gorm:"type:varchar(100)"`
	Username       string    `gorm:"type:varchar(100);unique"`
	Email          string    `gorm:"type:varchar(255);unique"`
	Password       string    `gorm:"type:varchar(255)"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Password != "" {
		hash, err := user.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = hash
	}

	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") && user.Password != "" {
		hash, err := user.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = hash
	}

	return
}

func (User) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
