package entity

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	CategoryID string   `json:"category_id" gorm:"type:uuid;not null"`
	Title      string   `json:"title" gorm:"type:varchar(255);not null"`
	Question   string   `json:"question" gorm:"type:text;not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	Answers    []Answer `json:"answers" gorm:"foreignKey:QuestionID"`
}

type QuestionOnly struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
