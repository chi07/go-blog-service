package model

type Paginator struct {
	CurrentPage int64
	Limit       int64
	Total       int64
	Offset      int64
}

type QueryString struct {
	Query string
}
