package response

type Pagination struct {
	Code        uint  `json:"code,omitempty"`
	Error       error `json:"error,omitempty"`
	TotalRecord int64 `json:"total_record,omitempty"`
	CurrentPage int   `json:"current_page,omitempty"`
	Limit       int   `json:"limit,omitempty"`
}
