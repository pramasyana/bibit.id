package shared

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type Empty struct{}

type JSONSchemaTemplate struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (t *JSONSchemaTemplate) SetData(data interface{}) {
	t.Data = data
}

func (t *JSONSchemaTemplate) SetFailed(code int, message string) *JSONSchemaTemplate {
	t.Code = code
	t.Message = message
	t.Success = false
	t.Data = Empty{}
	return t
}

func (t *JSONSchemaTemplate) SetSuccess(data interface{}, message string) *JSONSchemaTemplate {
	t.Code = http.StatusOK
	t.Message = message
	t.Success = true
	t.Data = data
	return t
}

func (t *JSONSchemaTemplate) ShowHTTPResponse(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(t.Code)
	return json.NewEncoder(c.Response()).Encode(t)
}
