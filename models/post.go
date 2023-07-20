package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"unique, not null"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
