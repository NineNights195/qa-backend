package usecase

import (
	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
)

type answerUsecase struct {
	answerRepo domain.AnswerRepo
}

func NewAnswerUsecase(answerRepo domain.AnswerRepo) domain.AnswerUseCase {
	return &answerUsecase{answerRepo: answerRepo}
}

func (u *answerUsecase) GetAnswersByQuestionID(questionID string) ([]entity.Answer, error) {
	return u.answerRepo.GetAnswersByQuestionID(questionID)
}

func (u *answerUsecase) CreateAnswer(answer *entity.Answer) error {
	return u.answerRepo.CreateAnswer(answer)
}

func (u *answerUsecase) UpdateAnswer(answer *entity.Answer) error {
	return u.answerRepo.UpdateAnswer(answer)
}

func (u *answerUsecase) DeleteAnswer(id string) error {
	return u.answerRepo.DeleteAnswer(id)
}
