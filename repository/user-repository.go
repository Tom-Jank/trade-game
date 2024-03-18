package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
    GetById(id uint) (*User, error)
}

type UserRepositoryImpl struct {
    gormDB *gorm.DB
}

type User struct {
    ID uint `gorm:"primaryKey"`
    Name string
}

func NewUserRepositoryImpl(gormDB *gorm.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{gormDB: gormDB}
}

func (ur *UserRepositoryImpl) GetById(id uint) (*User, error) {
    var user User
    if err := ur.gormDB.First(&user, id).Error; err != nil {
        return &user, err
    }
    return &user, nil
}
