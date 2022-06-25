package utils

import (
	"encoding/json"
	"evolvecredit-test/api/v1/dto"
	"evolvecredit-test/internal/common/pagination"
	"net/http"
)

func SendResponse(res http.ResponseWriter, status string, message string, data []interface{}, pagination_data pagination.PaginationData, statusCode int) {
	response := dto.ResponseDTO{}
	response.Status = status
	response.Message = message
	response.Data = data
	response.PaginationData = dto.GetPaginationDTOFromPaginationData(pagination_data)

	if statusCode == http.StatusInternalServerError {
		statusCode = http.StatusBadRequest
	}

	res.WriteHeader(statusCode)
	encoder := json.NewEncoder(res)
	encoder.Encode(response)
}
