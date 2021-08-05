package usecase

import (
	"github.com/pramasyana/bibit.id/src/modules/question/domain"
)

type QuestionUsecaseImpl struct{}

func NewQuestionUsecaseImpl() *QuestionUsecaseImpl {
	return &QuestionUsecaseImpl{}
}

func (q *QuestionUsecaseImpl) QuestionNumber1() (resp domain.ResponseNumber1, err error) {
	resp.Database = "Mysql"
	resp.Query = "SELECT id, username, (select username from `user` where id=us.parent) as parentUsername FROM `user` as us"
	return resp, nil
}
