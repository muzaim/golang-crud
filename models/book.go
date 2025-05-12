package models

import "time"

type BookResponse struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(100)" json:"title"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Price       int    `gorm:"type:integer" json:"price"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
}

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	Description string    `gorm:"type:varchar(100)" json:"description"`
	Price       int       `gorm:"type:integer" json:"price"`
	Author      string    `gorm:"type:varchar(100)" json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
