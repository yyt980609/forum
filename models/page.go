package models

type Page[T any] struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	Total       int64 `json:"total"`
	Pages       int64 `json:"pages"`
	Data        []T   `json:"data"`
}
