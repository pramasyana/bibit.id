package usecase

import "github.com/pramasyana/bibit.id/src/modules/question/domain"

type QuestionUsecase interface {
	QuestionNumber1() (resp domain.ResponseNumber1, err error)
}
