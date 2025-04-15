package domain

import "github.com/thanarat/qa-backend/entity"

type AnswerUseCase interface {
	GetAnswersByQuestionID(questionID string) ([]entity.Answer, error)
	CreateAnswer(answer *entity.Answer) error
	UpdateAnswer(answer *entity.Answer) error
	DeleteAnswer(id string) error
}

type AnswerRepo interface {
	GetAnswersByQuestionID(questionID string) ([]entity.Answer, error)
	CreateAnswer(answer *entity.Answer) error
	UpdateAnswer(answer *entity.Answer) error
	DeleteAnswer(id string) error
}
