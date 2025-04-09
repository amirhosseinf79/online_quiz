package dto

type PageFilter struct {
	Page     int `json:"page" query:"page" validate:"gte=0"`
	PageSize int `json:"page_size" query:"page_size" validate:"gte=0"`
}

func (p *PageFilter) SetDefaults() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
