package utils

import (
	"fmt"
	"strings"
)

type PaginationOptions struct {
	SelectColumns string
	BaseQuery     string
	SearchColumns []string
	OrderOptions  string
	Search        string
	Page          uint
	PerPage       uint
}

type PaginationQueryResult struct {
	Query      string
	CountQuery string
}

func GeneratePaginationQuery(options *PaginationOptions) *PaginationQueryResult {
	query := fmt.Sprintf("select %s %s", options.SelectColumns, options.BaseQuery)

	countQuery := fmt.Sprintf("select count(*) as total %s", options.BaseQuery)

	if options.Search != "" {
		searchClauses := make([]string, len(options.SearchColumns))
		for i, column := range options.SearchColumns {
			searchClauses[i] = fmt.Sprintf("%s like '%%%s%%'", column, options.Search)
		}
		searchQuery := strings.Join(searchClauses, " or ")

		if strings.Contains(query, " WHERE ") || strings.Contains(query, " where ") {
			query = fmt.Sprintf("%s and (%s)", query, searchQuery)
		} else {
			query = fmt.Sprintf("%s where (%s)", query, searchQuery)
		}

		if strings.Contains(countQuery, " WHERE ") || strings.Contains(countQuery, " where ") {
			countQuery = fmt.Sprintf("%s and (%s)", countQuery, searchQuery)
		} else {
			countQuery = fmt.Sprintf("%s where (%s)", countQuery, searchQuery)
		}
	}

	if options.OrderOptions != "" {
		query = fmt.Sprintf("%s order by %s", query, options.OrderOptions)
	}

	if options.Page != 0 && options.PerPage != 0 {
		offset := (options.Page - 1) * options.PerPage
		query = fmt.Sprintf("%s limit %d offset %d", query, options.PerPage, offset)
	}

	return &PaginationQueryResult{
		Query:      query,
		CountQuery: countQuery,
	}
}
