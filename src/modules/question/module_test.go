package question

import (
	"testing"
)

func TestNewModule(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Case 1: Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewModule()
		})
	}
}
