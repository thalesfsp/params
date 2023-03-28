//////
// Shared utils.
//////

package shared

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should work",
			args: args{
				v: map[string]interface{}{
					"a": 1,
					"b": 2,
				},
			},
			want:    []byte(`{"a":1,"b":2}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestProcess(t *testing.T) {
	t.Setenv("id2", "456")

	type Test struct {
		ID1 string `json:"id1" default:"123"`
		ID2 string `json:"id2" env:"id2"`
		ID3 string `json:"id3" id:"uuid"`
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Should work",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := Test{}

			if err := Process(&test); (err != nil) != tt.wantErr {
				t.Errorf("ProcessParams() error = %+v, wantErr %+v", err, tt.wantErr)
			}

			if test.ID1 != "123" {
				t.Errorf("ProcessParams() error = %+v, wantErr %+v", test.ID1, tt.wantErr)

				return
			}

			if test.ID2 != "456" {
				t.Errorf("ProcessParams() error = %+v, wantErr %+v", test.ID2, tt.wantErr)

				return
			}

			if len(test.ID3) != 36 {
				t.Errorf("ProcessParams() error = %+v, wantErr %+v", test.ID3, tt.wantErr)

				return
			}
		})
	}
}
