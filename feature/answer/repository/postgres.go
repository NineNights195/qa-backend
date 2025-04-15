package repository

import (
	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
	"gorm.io/gorm"
)

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepo(db *gorm.DB) domain.AnswerRepo {
	return &answerRepository{db: db}
}

func (r *answerRepository) GetAnswersByQuestionID(questionID string) ([]entity.Answer, error) {
	var answers []entity.Answer
	err := r.db.Where("question_id = ?", questionID).Find(&answers).Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *answerRepository) CreateAnswer(answer *entity.Answer) error {
	return r.db.Create(answer).Error
}

func (r *answerRepository) UpdateAnswer(answer *entity.Answer) error {
	return r.db.Save(answer).Error
}

func (r *answerRepository) DeleteAnswer(id string) error {
	return r.db.Delete(&entity.Answer{}, "id = ?", id).Error
}
