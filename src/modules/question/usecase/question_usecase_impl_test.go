package usecase

import (
	"testing"

	"github.com/pramasyana/bibit.id/src/modules/question/domain"
)

func TestNewQuestionUsecaseImpl(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Case 1: Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewQuestionUsecaseImpl()
		})
	}
}

func TestQuestionUsecaseImpl_QuestionNumber1(t *testing.T) {
	tests := []struct {
		name string
		q    *QuestionUsecaseImpl
	}{
		{
			name: "Case 1: Success",
			q:    NewQuestionUsecaseImpl(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuestionUsecaseImpl{}
			q.QuestionNumber1()
		})
	}
}

func TestQuestionUsecaseImpl_QuestionNumber3(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		q       *QuestionUsecaseImpl
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			q:    NewQuestionUsecaseImpl(),
			args: args{
				text: "ini adalah (Ryan Pramasyana)",
			},
			wantErr: false,
		},
		{
			name: "Case 2: Error (",
			q:    NewQuestionUsecaseImpl(),
			args: args{
				text: "ini adalah Ryan Pramasyana)",
			},
			wantErr: true,
		},
		{
			name: "Case 3: Error )",
			q:    NewQuestionUsecaseImpl(),
			args: args{
				text: "ini adalah (Ryan Pramasyana",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuestionUsecaseImpl{}
			_, err := q.QuestionNumber3(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionUsecaseImpl.QuestionNumber3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQuestionUsecaseImpl_QuestionNumber4(t *testing.T) {
	type args struct {
		text []string
	}
	tests := []struct {
		name    string
		q       *QuestionUsecaseImpl
		args    args
		wantErr bool
	}{
		{
			name: "case 1: Success",
			q:    NewQuestionUsecaseImpl(),
			args: args{
				text: []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuestionUsecaseImpl{}
			_, err := q.QuestionNumber4(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionUsecaseImpl.QuestionNumber4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQuestionUsecaseImpl_QuestionNumber2List(t *testing.T) {
	type args struct {
		params domain.ParamsNumber2
	}
	tests := []struct {
		name    string
		q       *QuestionUsecaseImpl
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			q:    NewQuestionUsecaseImpl(),
			args: args{
				params: domain.ParamsNumber2{
					Search: "batman",
					Page:   "1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuestionUsecaseImpl{}
			_, err := q.QuestionNumber2List(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionUsecaseImpl.QuestionNumber2List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQuestionUsecaseImpl_QuestionNumber2Detail(t *testing.T) {
	type args struct {
		params domain.ParamsNumber2
	}
	tests := []struct {
		name    string
		q       *QuestionUsecaseImpl
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			q:    NewQuestionUsecaseImpl(),
			args: args{
				params: domain.ParamsNumber2{
					ID: "tt0372784",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuestionUsecaseImpl{}
			_, err := q.QuestionNumber2Detail(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionUsecaseImpl.QuestionNumber2Detail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
