package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category  string     `json:"category"`
	Questions []Question `json:"questions" gorm:"foreignKey:CategoryID"`
}

type CategoryNoQuestion struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}
