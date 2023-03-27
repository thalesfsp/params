//////
// Params used by APIs, Storage, Adapters, etc.
//////

package param

import (
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
				t.Errorf("Delete.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.l.Soft != tt.want {
				t.Errorf("Delete.Validate() tt.l.Soft = %v, want %v", tt.l.Soft, tt.want)
			}
		})
	}
}
