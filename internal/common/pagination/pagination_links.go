package pagination

type PaginationData struct {
	FirstPage bool
	LastPage  bool
	PrevPage  int
	NextPage  int
	Total     int
}
