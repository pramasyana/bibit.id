package shared

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestJSONSchemaTemplate_SetData(t *testing.T) {
	type fields struct {
		Success bool
		Code    int
		Message string
		Data    interface{}
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Case 1: Success",
			args: args{
				data: `{}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &JSONSchemaTemplate{
				Success: tt.fields.Success,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Data:    tt.fields.Data,
			}
			tr.SetData(tt.args.data)
		})
	}
}

func TestJSONSchemaTemplate_SetFailed(t *testing.T) {
	type fields struct {
		Success bool
		Code    int
		Message string
		Data    interface{}
	}
	type args struct {
		code    int
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Case 1: Success",
			args: args{
				code:    http.StatusOK,
				message: "Success",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &JSONSchemaTemplate{
				Success: tt.fields.Success,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Data:    tt.fields.Data,
			}
			tr.SetFailed(tt.args.code, tt.args.message)
		})
	}
}

func TestJSONSchemaTemplate_SetSuccess(t *testing.T) {
	type fields struct {
		Success bool
		Code    int
		Message string
		Data    interface{}
	}
	type args struct {
		data    interface{}
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Case 1: Success",
			args: args{
				data:    `{}`,
				message: "Success",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &JSONSchemaTemplate{
				Success: tt.fields.Success,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Data:    tt.fields.Data,
			}
			tr.SetSuccess(tt.args.data, tt.args.message)
		})
	}
}

func TestJSONSchemaTemplate_ShowHTTPResponse(t *testing.T) {
	type fields struct {
		Success bool
		Code    int
		Message string
		Data    interface{}
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
			name: "Case 1; Success",
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
			tr := &JSONSchemaTemplate{
				Success: tt.fields.Success,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Data:    tt.fields.Data,
			}
			if err := tr.ShowHTTPResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("JSONSchemaTemplate.ShowHTTPResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
