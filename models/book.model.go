package models

import "time"

type BooksModel struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"book_id"`
	BookName   string    `gorm:"size:255;not null" json:"book_name"`
	BookPrice  float64   `gorm:"size:255;not null" json:"book_price"`
	BookAuthor string    `gorm:"size:255;not null" json:"book_author"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}
