package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID			uuid.UUID	`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email 		string		`gorm:"type:varchar(100);uniqueIndex;not null"`
	Password 	string		`gorm:"type:varchar(100);not null"`
	FullName	string		`gorm:"type:varchar(150);not null"`
	Active		bool
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}

type UserRepository interface {
	Insert(context.Context, User) error
	Update(context.Context, User) error
	Delete(context.Context, string) error
	FindById(context.Context, string) (User, error)
	FindByEmail(context.Context, string) (User, error)
	FindAll(context.Context) ([]User, error)
}