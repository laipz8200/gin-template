package model

import (
	"_template_/domain/entity"
	"_template_/domain/repository"
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     string  `gorm:"primaryKey"`
	Amount float64 `gorm:"default:0;not null"`
}

var _ repository.IUsers = (*userRepository)(nil)

type userRepository struct {
	conn *gorm.DB
}

// FindOne implements repository.IUsers
func (u *userRepository) FindOne(ctx context.Context, id string) (entity.User, error) {
	var user User
	if res := u.conn.WithContext(ctx).Where("id = ?", id).Find(&user); res.Error != nil {
		return entity.User{}, res.Error
	}

	return entity.NewUser(user.ID, user.Amount), nil
}

// Save implements repository.IUsers
func (u *userRepository) Save(ctx context.Context, user *entity.User) error {
	dbUser := User{
		ID:     user.ID(),
		Amount: user.Amount(),
	}

	if res := u.conn.WithContext(ctx).Save(&dbUser); res.Error != nil {
		return res.Error
	}

	return nil
}

func NewUserRepository(conn *gorm.DB) repository.IUsers {
	return &userRepository{
		conn: conn,
	}
}
