package param

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalesfsp/params/customsort"
)

func TestPagination_Process(t *testing.T) {
	tests := []struct {
		name    string
		p       *ListResult[string]
		wantErr bool
	}{
		{
			name: "valid pagination",
			p: &ListResult[string]{
				Meta: List{
					Count:  10,
					Fields: Fields{"name", "age"},
					Limit:  20,
					Offset: 0,
					Search: "test",
					Sort: customsort.SortMap{
						"name": customsort.Asc,
						"age":  customsort.Desc,
					},
				},
				Results: []string{"foo", "bar", "baz"},
			},
			wantErr: false,
		},
		{
			name: "invalid pagination",
			p: &ListResult[string]{
				Meta: List{
					Count:  -10,
					Fields: Fields{},
					Limit:  0,
					Offset: -1,
					Search: "",
					Sort: customsort.SortMap{
						"name": customsort.Asc,
						"age":  customsort.Desc,
					},
				},
				Results: []string{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.Process()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
