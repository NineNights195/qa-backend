package domain

import "github.com/thanarat/qa-backend/entity"

type CategoryUseCase interface {
	GetAllCategories() ([]entity.CategoryNoQuestion, error)
}

type CategoryRepo interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryById(id string) (entity.CategoryNoQuestion, error)
}
