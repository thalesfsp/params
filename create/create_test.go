package create

import (
	"reflect"
	"testing"
)

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
				t.Errorf("Set.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *Set
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name: "Should work",
			want: &Set{
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
				t.Errorf("Set.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}