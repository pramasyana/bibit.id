package usecase

import "github.com/pramasyana/bibit.id/src/modules/question/domain"

type QuestionUsecase interface {
	QuestionNumber1() (resp domain.ResponseNumber1, err error)
	QuestionNumber2List(text string) (resp domain.ResponseNumber3, err error)
	QuestionNumber2Detail(text string) (resp domain.ResponseNumber3, err error)
	QuestionNumber3(text string) (resp domain.ResponseNumber3, err error)
	QuestionNumber4(text []string) (resp domain.ResponseNumber4, err error)
}
