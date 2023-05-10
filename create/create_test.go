package create

import (
	"reflect"
	"testing"
)

func TestSet_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s       *Create
		wantErr bool
	}{
		{
			name:    "Should work",
			s:       &Create{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Create{}
			if err := s.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Create.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *Create
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name: "Should work",
			want: &Create{
				ID: "",
			},
			wantErr:           false,
			wantValidationErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %+v, want %+v", got, tt.want)
			}

			if err := got.Process(); (err != nil) != tt.wantValidationErr {
				t.Errorf("Create.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}
