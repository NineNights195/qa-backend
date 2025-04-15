package domain

import "github.com/thanarat/qa-backend/entity"

type QuestionUseCase interface {
	GetAllQuestions() ([]entity.QuestionOnly, error)
	GetQuestionByID(id string) (*entity.Question, error)
	CreateQuestion(question *entity.Question) error
	UpdateQuestion(question *entity.Question) error
	DeleteQuestion(id string) error
}

type QuestionRepo interface {
	GetAllQuestions() ([]entity.Question, error)
	GetQuestionByID(id string) (*entity.Question, error)
	CreateQuestion(question *entity.Question) error
	UpdateQuestion(question *entity.Question) error
	DeleteQuestion(id string) error
}
