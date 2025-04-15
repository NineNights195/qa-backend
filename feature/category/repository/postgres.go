package repository

import (
	"fmt"

	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) domain.CategoryRepo {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAllCategories() ([]entity.Category, error) {
	fmt.Println("[CategoryRepo.GetAllCategories]")
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(categories)
	return categories, nil
}

func (r *categoryRepository) GetCategoryById(id string) (entity.CategoryNoQuestion, error) {
	var category entity.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return entity.CategoryNoQuestion{}, err
	}

	return entity.CategoryNoQuestion{
		ID:       id,
		Category: category.Category,
	}, nil
}
