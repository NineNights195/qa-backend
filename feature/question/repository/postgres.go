package repository

import (
	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) domain.QuestionRepo {
	return &questionRepository{db: db}
}

func (r *questionRepository) GetAllQuestions() ([]entity.Question, error) {
	var questions []entity.Question
	err := r.db.Preload("Category").Preload("Answers").Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *questionRepository) GetQuestionByID(id string) (*entity.Question, error) {
	var question entity.Question
	err := r.db.Preload("Category").Preload("Answers").First(&question, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *questionRepository) CreateQuestion(question *entity.Question) error {
	return r.db.Create(question).Error
}

func (r *questionRepository) UpdateQuestion(question *entity.Question) error {
	return r.db.Save(question).Error
}

func (r *questionRepository) DeleteQuestion(id string) error {
	return r.db.Delete(&entity.Question{}, "id = ?", id).Error
}
