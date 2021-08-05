package delivery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/pramasyana/bibit.id/src/mock"
	"github.com/pramasyana/bibit.id/src/modules/question/usecase"
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
				QuestionUsecase: usecase.NewQuestionUsecaseImpl(),
			},
			args: args{
				c: echo.New().NewContext(
					httptest.NewRequest(http.MethodGet, "/number3", nil),
					httptest.NewRecorder(),
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := new(mock.SupplierUsecase)
			mock.On("QuestionNumber3", mock.Anything).Return(tt.wantErr)
			h := &EchoHandler{
				QuestionUsecase: tt.fields.QuestionUsecase,
			}
			if err := h.Number3(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("EchoHandler.Number3() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
