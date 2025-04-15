package usecase

import (
	"strconv"

	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
)

type questionUsecase struct {
	questionRepo domain.QuestionRepo
	categoryRepo domain.CategoryRepo
}

func NewQuestionUsecase(questionRepo domain.QuestionRepo, categoryRepo domain.CategoryRepo) domain.QuestionUseCase {
	return &questionUsecase{questionRepo: questionRepo, categoryRepo: categoryRepo}
}

func (u *questionUsecase) GetAllQuestions() ([]entity.QuestionOnly, error) {
	questions, err := u.questionRepo.GetAllQuestions()
	if err != nil {
		return nil, err
	}

	// Convert to QuestionOnly format
	questionOnly := make([]entity.QuestionOnly, len(questions))
	for i, q := range questions {
		category, err := u.categoryRepo.GetCategoryById(q.CategoryID)
		if err != nil {
			return nil, err
		}
		questionOnly[i] = entity.QuestionOnly{
			ID:        strconv.FormatUint(uint64(q.ID), 10),
			Category:  category.Category,
			Title:     q.Title,
			Question:  q.Question,
			CreatedAt: q.CreatedAt,
			UpdatedAt: q.UpdatedAt,
		}
	}

	return questionOnly, nil
}

func (u *questionUsecase) GetQuestionByID(id string) (*entity.Question, error) {
	return u.questionRepo.GetQuestionByID(id)
}

func (u *questionUsecase) CreateQuestion(question *entity.Question) error {
	return u.questionRepo.CreateQuestion(question)
}

func (u *questionUsecase) UpdateQuestion(question *entity.Question) error {
	return u.questionRepo.UpdateQuestion(question)
}

func (u *questionUsecase) DeleteQuestion(id string) error {
	return u.questionRepo.DeleteQuestion(id)
}
