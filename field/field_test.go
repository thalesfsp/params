package field

import (
	"testing"
)

func TestFields_String(t *testing.T) {
	tests := []struct {
		name string
		f    Fields
		want string
	}{
		{
			name: "Should work",
			f:    Fields{"a", "b", "c"},
			want: "a,b,c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("Fields.String() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestFields_UnmarshalParam(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		f       *Fields
		args    args
		wantErr bool
	}{
		{
			name:    "Should work",
			f:       &Fields{"a", "b", "c"},
			args:    args{src: "a,b,c"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.UnmarshalParam(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Fields.UnmarshalParam() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestFields_ToElasticSearch(t *testing.T) {
	tests := []struct {
		name    string
		f       *Fields
		want    string
		wantErr bool
	}{
		{
			name:    "Should work",
			f:       &Fields{"a", "b", "c"},
			want:    `["a","b","c"]`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.ToElasticSearch()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fields.ToElasticSearch() error = %+v, wantErr %+v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("Fields.ToElasticSearch() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
