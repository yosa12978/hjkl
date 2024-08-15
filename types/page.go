package types

type Page[T any] struct {
	Content     []T  `json:"content"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
	Total       int  `json:"total"`
	Size        int  `json:"size"`
}
