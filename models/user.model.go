package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32         `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string         `gorm:"size:255" json:"user_name"`
	Email     string         `gorm:"unique;size:255;not null" json:"email"`
	Password  string         `gorm:"size:255" json:"password"`
	Role      string         `gorm:"size:255" json:"role"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
