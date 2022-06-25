package dto

import "evolvecredit-test/internal/common/pagination"

type PaginationDTO struct {
	FirstPage bool `json:"first_page"`
	LastPage  bool `json:"last_page"`
	PrevPage  int  `json:"prev_page"`
	NextPage  int  `json:"next_page"`
	Total     int  `json:"total"`
}

func GetPaginationDTOFromPaginationData(data pagination.PaginationData) PaginationDTO {
	return PaginationDTO{
		FirstPage: data.FirstPage,
		LastPage:  data.LastPage,
		PrevPage:  data.PrevPage,
		NextPage:  data.NextPage,
		Total:     data.Total,
	}
}

type ResponseDTO struct {
	Status         string        `json:"status"`
	Message        string        `json:"message"`
	Data           []interface{} `json:"data"`
	PaginationData PaginationDTO `json:"pagination_data"`
}
