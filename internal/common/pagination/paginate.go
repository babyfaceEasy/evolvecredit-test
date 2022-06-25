package pagination

type Pagination struct {
	Page       int
	Limit      int
	TotalCount int
}

func NewPagination() Pagination {
	return Pagination{Page: 1, Limit: 10, TotalCount: 0}
}

func (p *Pagination) SetPage(page int) {
	if page <= 0 {
		page = 1
	}
	p.Page = page
}

func (p *Pagination) SetLimit(limit int) {
	if limit <= 0 {
		limit = 10
	}
	p.Limit = limit
}

func (p *Pagination) SetTotalCount(total int) {
	if total < 0 {
		total = 0
	}
	p.TotalCount = total
}

func (p *Pagination) getLastPage() bool {
	return p.Page * p.Limit >= p.TotalCount
}

func (p *Pagination) GetPaginationData() PaginationData {
	firstPage := true
	prevPage := 1
	if p.Page > 1 {
		firstPage = false
		prevPage = p.Page - 1
	}

	lastPage := p.getLastPage()
	nextPage := p.Page
	if !lastPage {
		nextPage = p.Page + 1
	}

	return PaginationData{
		FirstPage: firstPage,
		LastPage:  lastPage,
		PrevPage:  prevPage,
		NextPage:  nextPage,
		Total:     p.TotalCount,
	}

}
