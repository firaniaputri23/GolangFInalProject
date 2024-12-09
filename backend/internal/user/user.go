package user

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRes struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserRes struct {
	AccessToken string `json:"access_token"`
	ID          uint   `json:"id"`
	Username    string `json:"username"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)
}