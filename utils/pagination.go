package utils

type MetadataPagination struct {
	Total      uint `json:"total"`
	Page       uint `json:"page"`
	Limit      uint `json:"limit"`
	TotalPages uint `json:"total_pages"`
}

type Pagination[T any] struct {
	Data     []T                `json:"data"`
	Metadata MetadataPagination `json:"metadata"`
}

func GetPagination(page, limit *uint) (int, int, int) {
	const defaultPage, defaultLimit, maxLimit = 1, 10, 100

	// Handle page
	p := defaultPage
	if page != nil && *page > 0 {
		p = int(*page)
	}

	// Handle limit
	l := defaultLimit
	if limit != nil {
		if *limit == 0 {
			l = -1 // Special case: fetch all data
		} else if *limit > 0 && *limit <= maxLimit {
			l = int(*limit)
		} else {
			l = maxLimit
		}
	}

	offset := (p - 1) * l
	return l, p, offset
}
