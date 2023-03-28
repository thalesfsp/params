package update

import (
	"reflect"
	"testing"
)

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
				t.Errorf("Update.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *Update
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name:              "Should work",
			want:              &Update{},
			wantErr:           false,
			wantValidationErr: true,
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
				t.Errorf("Update.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}
