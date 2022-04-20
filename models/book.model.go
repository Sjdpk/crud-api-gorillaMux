package models

import (
	"time"

	"gorm.io/gorm"
)

type BooksModel struct {
	ID              uint32         `gorm:"primary_key;auto_increment" json:"book_id"`
	BookName        string         `gorm:"size:255;not null" json:"book_name"`
	BookPrice       float64        `gorm:"size:255;not null" json:"book_price"`
	BookAuthor      string         `gorm:"size:255;not null" json:"book_author"`
	BookDescription string         `gorm:"size:255;not null" json:"book_description"`
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type BookSubModel struct {
	ID         uint32    `json:"book_id"`
	BookName   string    `json:"book_name"`
	BookPrice  float64   `json:"book_price"`
	BookAuthor string    `json:"book_author"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
