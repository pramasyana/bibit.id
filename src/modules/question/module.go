package question

import (
	"github.com/pramasyana/bibit.id/src/modules/question/delivery"
	"github.com/pramasyana/bibit.id/src/modules/question/usecase"
)

type Module struct {
	RestHandler *delivery.EchoHandler
}

func NewModule() *Module {
	questionUsecase := usecase.NewQuestionUsecaseImpl()
	restHandler := delivery.NewEchoHandler(questionUsecase)
	return &Module{
		RestHandler: restHandler,
	}
}
