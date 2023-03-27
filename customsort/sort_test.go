//////
// Params used by APIs, Storage, Adapters, etc.
//////

package customsort

import (
	"reflect"
	"testing"
)

func TestSort_String(t *testing.T) {
	tests := []struct {
		name string
		s    Sort
		want string
	}{
		{
			name: "Should work",
			s:    "a:asc,b:desc",
			want: "a:asc,b:desc",
		},
		{
			name: "Should work",
			s:    "a:asc-b:desc",
			want: "a:asc-b:desc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Sort.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort_ToAnyString(t *testing.T) {
	type args struct {
		betweenKVSeparator      string
		betweenEntriesSeparator string
	}
	tests := []struct {
		name    string
		s       Sort
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should work",
			s:    "a:asc,b:desc",
			args: args{
				betweenKVSeparator:      " ",
				betweenEntriesSeparator: ",",
			},
			want:    "a asc,b desc",
			wantErr: false,
		},
		{
			name: "Should work",
			s:    "a:asc-b:desc",
			args: args{
				betweenKVSeparator:      " ",
				betweenEntriesSeparator: ",",
			},
			want:    "a asc,b desc",
			wantErr: false,
		},
		{
			name: "Should work",
			s:    "a=asc-b=desc",
			args: args{
				betweenKVSeparator:      " ",
				betweenEntriesSeparator: ",",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToAnyString(tt.args.betweenKVSeparator, tt.args.betweenEntriesSeparator)
			if got != tt.want {
				t.Errorf("Sort.ToAnyString() = %v, want %v", got, tt.want)
			}

			if err != nil && !tt.wantErr {
				t.Errorf("Sort.ToAnyString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSort_ToMap(t *testing.T) {
	tests := []struct {
		name    string
		s       Sort
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "Should work",
			s:       "a:asc,b:desc",
			want:    map[string]string{"a": Asc, "b": Desc},
			wantErr: false,
		},
		{
			name:    "Should fail",
			s:       "a=asc,b=desc",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Sort.ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortMap_UnmarshalParam(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		sM      SortMap
		args    args
		wantErr bool
	}{
		{
			name: "Should work",
			sM:   SortMap{},
			args: args{
				src: "a:asc,b:desc",
			},
			wantErr: false,
		},
		{
			name: "Should fail",
			sM:   SortMap{},
			args: args{
				src: "a=asc,b=desc",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sM.UnmarshalParam(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortMap.UnmarshalParam() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if tt.sM["a"] != Asc {
					t.Errorf("SortMap.UnmarshalParam() error = %v, wantErr %v", err, tt.wantErr)
				}

				if tt.sM["b"] != Desc {
					t.Errorf("SortMap.UnmarshalParam() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestNewFromMap(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want Sort
	}{
		{
			name: "Should work",
			args: args{
				m: map[string]string{
					"a": Asc,
					"b": Desc,
				},
			},
			want: "a:asc,b:desc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromMap(tt.args.m); got != tt.want {
				t.Errorf("NewFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort_Validate(t *testing.T) {
	tests := []struct {
		name string
		s    Sort
		want bool
	}{
		{
			name: "Should work",
			s:    "a:asc,b:desc",
			want: true,
		},
		{
			name: "Should fail",
			s:    "a=asc,b=desc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValid(); got != tt.want {
				t.Errorf("Sort.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortMap_ToSort(t *testing.T) {
	tests := []struct {
		name string
		sM   *SortMap
		want Sort
	}{
		{
			name: "Should work",
			sM: &SortMap{
				"a": Asc,
				"b": Desc,
			},
			want: "a:asc,b:desc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sM.ToSort(); got != tt.want {
				t.Errorf("SortMap.ToSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
