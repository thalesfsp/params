package param

import (
	"testing"
)

func TestCount_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       *Count
		wantErr bool
	}{
		{
			name: "Should work",
			c: &Count{
				Search: "test",
			},
			wantErr: false,
		},
		{
			name: "Should fail - negative limit",
			c: &Count{
				Search: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Count.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
