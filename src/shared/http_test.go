package shared

import (
	"io"
	"net/http"
	"testing"
)

func TestRequestHTTP(t *testing.T) {
	type args struct {
		method  string
		url     string
		body    io.Reader
		target  interface{}
		headers []map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Case 1: Success",
			args: args{
				method:  http.MethodGet,
				url:     "http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=1",
				body:    nil,
				target:  nil,
				headers: nil,
			},
			want:    http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := RequestHTTP(tt.args.method, tt.args.url, tt.args.body, tt.args.target, tt.args.headers...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
