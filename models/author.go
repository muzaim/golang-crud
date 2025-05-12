package models

import "time"

type AuthorResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
}

type Author struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Gender    string    `gorm:"type.char(1)" json:"gender"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Age       int       `gorm:"type:integer" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
