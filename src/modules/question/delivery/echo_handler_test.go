package delivery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	mocksQuestionUsecase "github.com/pramasyana/bibit.id/mocks/src/modules/question/usecase"
	"github.com/pramasyana/bibit.id/src/modules/question/domain"
	"github.com/pramasyana/bibit.id/src/modules/question/usecase"
	"github.com/stretchr/testify/mock"
)

func TestNewEchoHandler(t *testing.T) {
	type args struct {
		questionUsecase usecase.QuestionUsecase
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Case 1: Success",
			args: args{
				questionUsecase: usecase.NewQuestionUsecaseImpl(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewEchoHandler(tt.args.questionUsecase)
		})
	}
}

func TestEchoHandler_Mount(t *testing.T) {
	type fields struct {
		QuestionUsecase usecase.QuestionUsecase
	}
	type args struct {
		group *echo.Group
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Case 1: Success",
			fields: fields{
				QuestionUsecase: usecase.NewQuestionUsecaseImpl(),
			},
			args: args{
				group: echo.New().Group("question"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				QuestionUsecase: tt.fields.QuestionUsecase,
			}
			h.Mount(tt.args.group)
		})
	}
}

func TestEchoHandler_Number1(t *testing.T) {
	type fields struct {
		QuestionUsecase usecase.QuestionUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			fields: fields{
				QuestionUsecase: usecase.NewQuestionUsecaseImpl(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodGet, "/number1", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				QuestionUsecase: tt.fields.QuestionUsecase,
			}
			if err := h.Number1(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("EchoHandler.Number1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEchoHandler_Number3(t *testing.T) {
	type fields struct {
		QuestionUsecase usecase.QuestionUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			fields: fields{
				QuestionUsecase: func() usecase.QuestionUsecase {
					mocksQuestionUsecase := new(mocksQuestionUsecase.QuestionUsecase)
					mocksQuestionUsecase.On("QuestionNumber3", mock.Anything).Return(domain.ResponseNumber3{}, nil)
					return mocksQuestionUsecase
				}(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodGet, "/number3", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
		{
			name: "Case 2: Error QuestionNumber3",
			fields: fields{
				QuestionUsecase: func() usecase.QuestionUsecase {
					mocksQuestionUsecase := new(mocksQuestionUsecase.QuestionUsecase)
					mocksQuestionUsecase.On("QuestionNumber3", mock.Anything).Return(domain.ResponseNumber3{}, errors.New("error"))
					return mocksQuestionUsecase
				}(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodGet, "/number3", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
		{
			name: "Case 3: Error Params",
			fields: fields{
				QuestionUsecase: func() usecase.QuestionUsecase {
					mocksQuestionUsecase := new(mocksQuestionUsecase.QuestionUsecase)
					mocksQuestionUsecase.On("QuestionNumber3", mock.Anything).Return(domain.ResponseNumber3{}, errors.New("error"))
					return mocksQuestionUsecase
				}(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodPost, "/number3", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				QuestionUsecase: tt.fields.QuestionUsecase,
			}
			if err := h.Number3(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("EchoHandler.Number3() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEchoHandler_Number4(t *testing.T) {
	type fields struct {
		QuestionUsecase usecase.QuestionUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			fields: fields{
				QuestionUsecase: func() usecase.QuestionUsecase {
					mocksQuestionUsecase := new(mocksQuestionUsecase.QuestionUsecase)
					mocksQuestionUsecase.On("QuestionNumber4", mock.Anything).Return(domain.ResponseNumber4{}, nil)
					return mocksQuestionUsecase
				}(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodGet, "/number3", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
		{
			name: "Case 2: Error QuestionNumber4",
			fields: fields{
				QuestionUsecase: func() usecase.QuestionUsecase {
					mocksQuestionUsecase := new(mocksQuestionUsecase.QuestionUsecase)
					mocksQuestionUsecase.On("QuestionNumber4", mock.Anything).Return(domain.ResponseNumber4{}, errors.New("error"))
					return mocksQuestionUsecase
				}(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodGet, "/number4", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
		{
			name: "Case 3: Error Params",
			fields: fields{
				QuestionUsecase: func() usecase.QuestionUsecase {
					mocksQuestionUsecase := new(mocksQuestionUsecase.QuestionUsecase)
					mocksQuestionUsecase.On("QuestionNumber4", mock.Anything).Return(domain.ResponseNumber4{}, errors.New("error"))
					return mocksQuestionUsecase
				}(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodPost, "/number4", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				QuestionUsecase: tt.fields.QuestionUsecase,
			}
			if err := h.Number4(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("EchoHandler.Number4() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
