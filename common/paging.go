package common

import "strings"

type Paging struct {
	Page  int64 `json:"page" form:"page"`
	Limit int64 `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
	//Suport cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}

	p.NextCursor = strings.TrimSpace(p.NextCursor)
}
