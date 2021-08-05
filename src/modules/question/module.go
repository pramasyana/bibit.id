package question

import "github.com/pramasyana/bibit.id/src/modules/question/delivery"

type Module struct {
	RestHandler *delivery.EchoHandler
}

func NewModule() *Module {
	restHandler := delivery.NewEchoHandler()
	return &Module{
		RestHandler: restHandler,
	}
}
