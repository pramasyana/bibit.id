package usecase

import (
	"errors"
	"sort"
	"strings"

	"github.com/pramasyana/bibit.id/src/modules/question/domain"
	"github.com/pramasyana/bibit.id/src/shared"
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

func (q *QuestionUsecaseImpl) QuestionNumber3(text string) (resp domain.ResponseNumber3, err error) {
	resp.Text = text
	i := strings.IndexByte(text, '(')
	if i < 0 {
		resp.Result = ""
		return resp, errors.New("Not Found Bucket On Text")
	}
	i++
	j := strings.IndexByte(text[i:], ')')
	if j < 0 {
		resp.Result = ""
		return resp, errors.New("Not Found Bucket On Text")
	}

	resp.Result = text[i : i+j]
	return resp, nil
}

func (q *QuestionUsecaseImpl) QuestionNumber4(text []string) (resp domain.ResponseNumber4, err error) {
	resp.Text = text
	var result [][]string
	newText := q.RemoveDuplicate(text)
	for _, val := range newText {
		var subResult []string
		sortWord1 := q.SortString(val)
		for _, valWord := range text {
			sortWord2 := q.SortString(valWord)
			if sortWord1 == sortWord2 {
				subResult = append(subResult, valWord)
			}
		}
		result = append(result, subResult)
	}
	resp.Result = result
	return resp, nil
}

func (q *QuestionUsecaseImpl) SortString(text string) string {
	str := strings.Split(text, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func (q *QuestionUsecaseImpl) RemoveDuplicate(text []string) []string {
	var tempkey = map[string]bool{}
	result := []string{}
	for _, val := range text {
		if _, value := tempkey[q.SortString(val)]; !value {
			tempkey[q.SortString(val)] = true
			result = append(result, q.SortString(val))
		}
	}
	return result
}

func (q *QuestionUsecaseImpl) QuestionNumber2List(params domain.ParamsNumber2) (resp *domain.ResponseNumber2List, err error) {
	response := &domain.ResponseNumber2List{}
	_, err = shared.RequestHTTP("GET", "http://www.omdbapi.com/?apikey=faf7e5bb&s="+params.Search+"&page="+params.Page, nil, response, nil)
	return response, err
}

func (q *QuestionUsecaseImpl) QuestionNumber2Detail(params domain.ParamsNumber2) (resp *domain.ResponseNumber2Detail, err error) {
	response := &domain.ResponseNumber2Detail{}
	_, err = shared.RequestHTTP("GET", "http://www.omdbapi.com/?apikey=faf7e5bb&i="+params.ID, nil, response, nil)
	return response, err
}
