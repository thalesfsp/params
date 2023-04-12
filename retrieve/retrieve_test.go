package retrieve

import (
	"reflect"
	"testing"
)

// DocumentID is the document ID.
const DocumentID = "VFzrpYMBXu5BQSZxo0qX"

func TestGet_Validate(t *testing.T) {
	tests := []struct {
		name    string
		g       *Retrieve
		wantErr bool
	}{
		{
			name: "Should work",
			g: &Retrieve{
				ID: DocumentID,
			},
			wantErr: false,
		},
		{
			name:    "Should fail - missing ID",
			g:       &Retrieve{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Get.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *Retrieve
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name: "Should work",
			want: &Retrieve{
				ID: "",
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
				t.Errorf("Get.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}
