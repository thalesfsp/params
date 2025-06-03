package listresult

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalesfsp/params/v2/customsort"
	"github.com/thalesfsp/params/v2/field"
	"github.com/thalesfsp/params/v2/list"
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
				Meta: list.List{
					Count:  10,
					Fields: field.Fields{"name", "age"},
					Limit:  20,
					Offset: 0,
					Search: "test",
					Sort:   [][]string{{"name", customsort.Asc}, {"age", customsort.Desc}},
				},
				Results: []string{"foo", "bar", "baz"},
			},
			wantErr: false,
		},
		{
			name: "invalid pagination",
			p: &ListResult[string]{
				Meta: list.List{
					Count:  -10,
					Fields: field.Fields{},
					Limit:  0,
					Offset: -1,
					Search: "",
					Sort:   [][]string{{"name", customsort.Asc}, {"age", customsort.Desc}},
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
