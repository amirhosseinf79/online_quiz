package dto

type PageFilter struct {
	Page     int `json:"page" query:"page" validate:"gte=0"`
	PageSize int `json:"page_size" query:"page_size" validate:"gte=0"`
}
