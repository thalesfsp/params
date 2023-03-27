package param

import "testing"

func TestSet_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s       *Set
		wantErr bool
	}{
		{
			name:    "Should work",
			s:       &Set{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{}
			if err := s.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Set.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
