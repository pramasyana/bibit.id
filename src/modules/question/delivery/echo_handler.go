package delivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pramasyana/bibit.id/src/modules/question/domain"
	"github.com/pramasyana/bibit.id/src/modules/question/usecase"
	"github.com/pramasyana/bibit.id/src/shared"
)

type EchoHandler struct {
	QuestionUsecase usecase.QuestionUsecase
}

func NewEchoHandler(questionUsecase usecase.QuestionUsecase) *EchoHandler {
	return &EchoHandler{
		QuestionUsecase: questionUsecase,
	}
}

func (h *EchoHandler) Mount(group *echo.Group) {
	group.GET("/number1", h.Number1)
	group.POST("/number3", h.Number3)
	group.POST("/number4", h.Number4)
}

// Number1 ...
func (h *EchoHandler) Number1(c echo.Context) error {
	response := new(shared.JSONSchemaTemplate)

	resp, _ := h.QuestionUsecase.QuestionNumber1()

	return response.SetSuccess(resp, "Success").ShowHTTPResponse(c)
}

// Number3 ...
func (h *EchoHandler) Number3(c echo.Context) error {
	response := new(shared.JSONSchemaTemplate)

	params := new(domain.ParamsNumber3)
	if err := c.Bind(params); err != nil {
		return response.SetFailed(http.StatusBadRequest, err.Error()).ShowHTTPResponse(c)
	}

	resp, err := h.QuestionUsecase.QuestionNumber3(params.Text)
	if err != nil {
		return response.SetFailed(http.StatusBadRequest, err.Error()).ShowHTTPResponse(c)
	}

	return response.SetSuccess(resp, "Success").ShowHTTPResponse(c)
}

// Number4 ...
func (h *EchoHandler) Number4(c echo.Context) error {
	response := new(shared.JSONSchemaTemplate)

	params := new(domain.ParamsNumber4)
	if err := c.Bind(params); err != nil {
		return response.SetFailed(http.StatusBadRequest, err.Error()).ShowHTTPResponse(c)
	}

	resp, err := h.QuestionUsecase.QuestionNumber4(params.Text)
	if err != nil {
		return response.SetFailed(http.StatusBadRequest, err.Error()).ShowHTTPResponse(c)
	}

	return response.SetSuccess(resp, "Success").ShowHTTPResponse(c)
}
