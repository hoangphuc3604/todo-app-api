package common

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (c *Paging) Process() {
	if c.Page <= 0 {
		c.Page = 1
	}

	if c.Limit <= 0 {
		c.Limit = 10
	}

	if c.Limit > 100 {
		c.Limit = 10
	}

	if c.Total <= 0 {
		c.Total = 0
	}
}

func (c *Paging) Offset() int {
	return (c.Page - 1) * c.Limit
}