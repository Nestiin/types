package types

type Pagination struct {
	Page          int           `json:"page"`
	NumElements   int           `json:"numElements"`
	PageSize      int           `json:"pageSize"`
	TotalPages    int           `json:"totalPages"`
	TotalElements int           `json:"totalElements"`
	Rows          []interface{} `json:"rows"`
}
