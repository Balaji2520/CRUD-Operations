package pagination

type Paginatevalue struct {
	Filters    string `json:"filters,omitempty" query:"filters"`
	Per_page   int    `json:"per_page,omitempty" query:"per_page"`
	Page_no    int    `json:"page_no,omitempty" query:"page_no"`
	Sort       string `json:"sort,omitempty" query:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

var BasePaginatevalue = Paginatevalue{Filters: "", Per_page: 10, Page_no: 1, Sort: "", TotalRows: 0, TotalPages: 0}

func (p *Paginatevalue) GetPage() int {
	if p.Page_no == 0 {
		p.Page_no = 1
	}
	return p.Page_no
}

func (p *Paginatevalue) GetOffset() int {
	return (p.GetPage() - 1) * p.Per_page
}

func (p *Paginatevalue) GetLimit() int {
	if p.Per_page == 0 {
		p.Per_page = 10
	}
	return p.Per_page
}

func (p *Paginatevalue) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
