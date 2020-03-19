package gfcommon

import "math"

type Page struct {
	Total     int         `form:"total" json:"total"`
	Limit     int         `form:"limit" json:"limit"`
	Page      int         `form:"page" json:"page"`
	TotalPage int         `json:"totalPage"`
	NextPage  int         `json:"page"`
	PrevPage  int         `json:"page"`
	HasNext   bool        `json:"page"`
	HasPrev   bool        `json:"page"`
	Data      interface{} `json:"data"`
}

func (page *Page) Paginate(total int, data interface{}) {
	page.Total = total
	page.Data = data
	page.TotalPage = int(math.Ceil(float64(total) / float64(page.Limit)))
	page.PrevPage = page.Page - 1
	page.HasPrev = true

	if page.PrevPage < 1 {
		page.PrevPage = 0
		page.HasPrev = false
	}
	page.NextPage = page.Page + 1
	page.HasNext = true
	if page.NextPage >= page.TotalPage {
		page.NextPage = 0
		page.HasNext = false
	}
}
