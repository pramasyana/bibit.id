package usecase

import "github.com/pramasyana/bibit.id/src/modules/question/domain"

type QuestionUsecase interface {
	QuestionNumber1() (resp domain.ResponseNumber1, err error)
	QuestionNumber2List(params domain.ParamsNumber2) (resp *domain.ResponseNumber2List, err error)
	QuestionNumber2Detail(params domain.ParamsNumber2) (resp *domain.ResponseNumber2Detail, err error)
	QuestionNumber3(text string) (resp domain.ResponseNumber3, err error)
	QuestionNumber4(text []string) (resp domain.ResponseNumber4, err error)
}
