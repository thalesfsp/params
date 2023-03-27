package param

import (
	"testing"
)

func TestGet_Validate(t *testing.T) {
	tests := []struct {
		name    string
		g       *Get
		wantErr bool
	}{
		{
			name: "Should work",
			g: &Get{
				ID: DocumentID,
			},
			wantErr: false,
		},
		{
			name:    "Should fail - missing ID",
			g:       &Get{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Get.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
