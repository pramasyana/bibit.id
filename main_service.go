package main

import "github.com/labstack/echo"

type Service struct {
	httpServer *echo.Echo
}

func InitServer() *Service {
	echoServer := echo.New()

	return &Service{
		httpServer: echoServer,
	}
}
