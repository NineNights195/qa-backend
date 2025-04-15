package usecase

import (
	"fmt"
	"strconv"

	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
)

type categoryUsecase struct {
	categoryRepo domain.CategoryRepo
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepo) domain.CategoryUseCase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

func (u *categoryUsecase) GetAllCategories() ([]entity.CategoryNoQuestion, error) {
	fmt.Println("[CategoryUsecase.GetAllCategories]")
	categories, err := u.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	// Convert to CategoryNoQuestion format
	categoryNoQuestions := make([]entity.CategoryNoQuestion, len(categories))
	for i, c := range categories {
		categoryNoQuestions[i] = entity.CategoryNoQuestion{
			ID:       strconv.FormatUint(uint64(c.ID), 10),
			Category: c.Category,
		}
	}

	return categoryNoQuestions, nil
}
