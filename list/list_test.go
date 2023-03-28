package list

import (
	"reflect"
	"testing"

	"github.com/thalesfsp/params/customsort"
)

func TestListParams_Validate(t *testing.T) {
	tests := []struct {
		name    string
		list    *List
		wantErr bool
	}{
		{
			name: "Should work",
			list: &List{
				Fields: []string{"a", "b"},
				Limit:  1,
				Offset: 1,
				Search: "ab",
				Sort: customsort.SortMap{
					"a": customsort.Asc,
					"b": customsort.Desc,
				},
			},
			wantErr: false,
		},
		{
			name: "Should fail - Fields - min 1",
			list: &List{
				Fields: []string{},
			},
			wantErr: true,
		},
		{
			name: "Should fail - Limit - negative",
			list: &List{
				Limit: -1,
			},
			wantErr: true,
		},
		{
			name: "Should fail - Offset - negative",
			list: &List{
				Offset: -1,
			},
			wantErr: true,
		},
		{
			name: "Should fail - Sort - invalid order",
			list: &List{
				Sort: customsort.SortMap{
					"a": "asd",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.list.Sort == nil {
				if !customsort.NewFromMap(tt.list.Sort).IsValid() && !tt.wantErr {
					t.Errorf("List.Sort.Validate() = %+v, want %+v", false, tt.wantErr)

					return
				}
			}

			if err := tt.list.Process(); (err != nil) != tt.wantErr && tt.list.Sort == nil {
				t.Errorf("List.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		want              *List
		wantErr           bool
		wantValidationErr bool
	}{
		{
			name: "Should work",
			want: &List{
				Limit: 10,
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
				t.Errorf("List.Process() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}
