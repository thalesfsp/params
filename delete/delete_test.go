package delete

import (
	"reflect"
	"testing"
)

// DocumentID is the document ID.
const DocumentID = "VFzrpYMBXu5BQSZxo0qX"

func TestDelete_Validate(t *testing.T) {
	tests := []struct {
		name    string
		l       Delete
		wantErr bool
		want    bool
	}{
		{
			name: "Should work",
			l: Delete{
				ID:   DocumentID,
				Soft: true,
			},
			wantErr: false,
			want:    true,
		},
		{
			name:    "Should fail - missing ID",
			l:       Delete{},
			wantErr: true,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Delete.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.l.Soft != tt.want {
				t.Errorf("Delete.Validate() tt.l.Soft = %+v, want %+v", tt.l.Soft, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *Delete
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name: "Should work",
			want: &Delete{
				ID:   "",
				Soft: false,
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
				t.Errorf("Delete.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}
