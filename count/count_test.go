package count

import (
	"reflect"
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
				t.Errorf("Count.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *Count
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name: "Should work",
			want: &Count{
				Search: "",
			},
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
				t.Errorf("Count.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}
