package utils

type PaginationResult[T any] struct {
	Data       []T
	Total      uint
	Page       uint
	PerPage    uint
	PageCounts uint
}
