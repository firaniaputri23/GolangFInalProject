package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
    Profile  string `json:"profile"`
    Friends  []User `gorm:"many2many:user_friends;"`
}
