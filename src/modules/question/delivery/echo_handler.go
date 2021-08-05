package delivery

import (
	"github.com/labstack/echo"
	"github.com/pramasyana/bibit.id/src/shared"
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
	response := new(shared.JSONSchemaTemplate)

	return response.SetSuccess(shared.Empty{}, "Success").ShowHTTPResponse(c)
}
