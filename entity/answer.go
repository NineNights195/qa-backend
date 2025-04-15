package entity

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	QuestionID string `json:"question_id" gorm:"type:uuid;not null"`
	Answer     string `json:"answer" gorm:"type:text;not null"`
}
