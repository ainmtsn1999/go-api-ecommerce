package models

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
}
