package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepsitory struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.Debug().AutoMigrate(&User{})
	return userRepsitory{db}
}

func (r userRepsitory) Insert(ctx context.Context, user User) error {
	user.ID = uuid.New()
	return r.db.Debug().WithContext(ctx).Create(&user).Error
}

func (r userRepsitory) Update(ctx context.Context, user User) error {	
	return r.db.Debug().WithContext(ctx).Where("id=?", user.ID).Updates(&user).Error
}

func (r userRepsitory) Delete(ctx context.Context, id string) error {
	return r.db.Debug().WithContext(ctx).Where("id=?", id).Delete(&User{}).Error
}

func (r userRepsitory) FindById(ctx context.Context, id string) (user User, err error) {
	err = r.db.Debug().WithContext(ctx).Where("id=?", id).First(&user).Error
	return user, err
}
func (r userRepsitory) FindByEmail(ctx context.Context, email string) (User, error) {
	var user = User{}
	err := r.db.Debug().WithContext(ctx).Find(&user, "email=?", email).Error
	return user, err
}

func (r userRepsitory) FindAll(ctx context.Context) (users []User, err error) {
	err = r.db.Debug().WithContext(ctx).Find(&users).Error
	return users, err
}