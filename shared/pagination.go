package shared

type Pagination struct {
	Page  int
	Limit int
}

func (p *Pagination) Normalize(defaultLimit int) {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

type PaginatedResponse[T any] struct {
	Data        []T  `json:"data"`
	Page        int  `json:"page"`
	Limit       int  `json:"limit"`
	TotalItems  int  `json:"totalItems"`
	TotalPages  int  `json:"totalPages"`
	HasNextPage bool `json:"hasNextPage"`
}

func NewPaginatedResponse[T any](data []T, page, limit, total int) PaginatedResponse[T] {
	totalPages := (total + limit - 1) / limit
	return PaginatedResponse[T]{
		Data:        data,
		Page:        page,
		Limit:       limit,
		TotalItems:  total,
		TotalPages:  totalPages,
		HasNextPage: page < totalPages,
	}
}
