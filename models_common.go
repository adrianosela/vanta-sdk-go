package vanta

type PageInfo struct {
	HasPreviousPage bool   `json:"has_previous_page"`
	HasNextPage     bool   `json:"has_next_page"`
	StartCursor     string `json:"start_cursor"`
	EndCursor       string `json:"end_cursor"`
}

type GenericNamedItem struct {
	Name string `json:"name"`
}
