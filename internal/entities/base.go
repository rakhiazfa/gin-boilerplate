package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationEntity struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-"`
}

type ApplicationEntityWithSoftDelete struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-"`
	DeletedAt gorm.DeletedAt
}
