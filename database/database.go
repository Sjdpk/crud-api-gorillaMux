package database

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID         uint           `gorm:"primaryKey"`
	BookName   string         `gorm:"size:255;not null" json:"book_name"`
	BookPrice  float64        `gorm:"size:255;not null" json:"book_price"`
	BookAuthor string         `gorm:"size:255;not null" json:"book_author"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}
