package repository

import (
    "github.com/<username>/<project>/internal/models"
    "gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user models.User) error
    FindUserByUsername(username string) (models.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(user models.User) error {
    return r.db.Create(&user).Error
}

func (r *userRepository) FindUserByUsername(username string) (models.User, error) {
    var user models.User
    err := r.db.Where("username = ?", username).First(&user).Error
    return user, err
}
