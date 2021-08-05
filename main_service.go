package main

import (
	"github.com/labstack/echo"
	"github.com/pramasyana/bibit.id/src/modules/question"
)

type Service struct {
	httpServer     *echo.Echo
	QuestionModule *question.Module
}

func InitServer() *Service {
	echoServer := echo.New()
	questionModule := question.NewModule()

	return &Service{
		httpServer:     echoServer,
		QuestionModule: questionModule,
	}
}
