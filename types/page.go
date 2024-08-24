package types

type Page[T any] struct {
	Content []T          `json:"content"`
	Meta    PageMetadata `json:"meta"`
}

type PageMetadata struct {
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
	Total       int  `json:"total"`
	Size        int  `json:"size"`
}
