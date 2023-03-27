package param

import "testing"

func TestUpdate_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u       *Update
		wantErr bool
	}{
		{
			name: "Should work",
			u: &Update{
				ID: "asdadsadas232",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Update.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
