package utils

import (
	"testing"
)

func TestGeneratePaginationQuery(t *testing.T) {
	tests := []struct {
		name    string
		options PaginationOptions
		want    PaginationQueryResult
	}{
		{
			name: "Basic Pagination",
			options: PaginationOptions{
				SelectColumns: "*",
				BaseQuery:     "from users",
				OrderOptions:  "id DESC",
				Page:          1,
				PerPage:       10,
			},
			want: PaginationQueryResult{
				Query:      "select * from users order by id DESC limit 10 offset 0",
				CountQuery: "select count(*) as total from users",
			},
		},
		// Add more test cases here for different scenarios
		{
			name: "With Search",
			options: PaginationOptions{
				SelectColumns: "*",
				BaseQuery:     "from users",
				SearchColumns: []string{"name", "email"},
				Search:        "John",
				Page:          1,
				PerPage:       10,
			},
			want: PaginationQueryResult{
				Query:      "select * from users where (name like '%John%' or email like '%John%') limit 10 offset 0",
				CountQuery: "select count(*) as total from users where (name like '%John%' or email like '%John%')",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GeneratePaginationQuery(&tt.options)
			if got.Query != tt.want.Query {
				t.Errorf("GeneratePaginationQuery().Query = %v, want %v", got.Query, tt.want.Query)
			}
			if got.CountQuery != tt.want.CountQuery {
				t.Errorf("GeneratePaginationQuery().CountQuery = %v, want %v", got.CountQuery, tt.want.CountQuery)
			}
		})
	}
}
