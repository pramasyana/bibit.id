package delivery

import (
	"github.com/labstack/echo"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Mount(group *echo.Group) {
	group.GET("/number1", h.Number1)
}

// Number1 ...
func (h *EchoHandler) Number1(c echo.Context) error {
	return nil
}
