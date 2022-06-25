package handlers

import (
	"errors"
	"evolvecredit-test/api/v1/dto"
	"evolvecredit-test/api/v1/messages"
	"evolvecredit-test/api/v1/utils"
	"evolvecredit-test/internal/common/pagination"
	"evolvecredit-test/internal/user"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

type UserHandler struct {
	userService user.UserService
	pagination  pagination.Pagination
}

func NewUserHandler(dbConn *bun.DB) (*UserHandler, error) {
	userRepository, err := user.NewUserRepository(dbConn)
	if err != nil {
		log.Println("Error occurred in building UserHandler: %w", err)
		return &UserHandler{}, err
	}

	userService, err := user.NewUserService(userRepository)
	if err != nil {
		log.Println("Error occurred in building UserHandler: %w", err)
		return &UserHandler{}, err
	}

	pagination := pagination.NewPagination()

	return &UserHandler{userService: *userService, pagination: pagination}, nil
}

func (h *UserHandler) ListAll(res http.ResponseWriter, req *http.Request) {
	page := strings.Trim(req.URL.Query().Get("page_no"), " ")
	limit := strings.Trim(req.URL.Query().Get("limit"), " ")

	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			log.Printf("error in converting page value :%s to int\n", page)
			pageInt = 1
		}
		h.pagination.SetPage(pageInt)
	}

	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			log.Printf("error in converting limit value :%s to int\n", limit)
			limitInt = 10
		}
		h.pagination.SetLimit(limitInt)
	}

	data := make([]interface{}, 0)
	log.Printf("pagination dt: %v\n", h.pagination)
	users, err := h.userService.List(h.pagination.Page, h.pagination.Limit)
	if err != nil {
		log.Printf("An error occurred while listing users: %v", err)
		pgData := h.pagination.GetPaginationData()
		utils.SendResponse(res, "failed", messages.ErrorOccurred, data, pgData, http.StatusInternalServerError)
		return

	}
	totalCount, err := h.userService.ListTotalCount()
	if err != nil {
		log.Printf("An error occurred while listing users: %v", err)
		utils.SendResponse(res, "failed", messages.ErrorOccurred, data,h.pagination.GetPaginationData(), http.StatusInternalServerError)
		return
	}

	h.pagination.SetTotalCount(totalCount)
	paginationData := h.pagination.GetPaginationData()


	for _, user := range users {
		data = append(data, dto.GetUserDTOFromUser(user))
	}
	utils.SendResponse(res, "success", "List of users", data, paginationData, http.StatusOK)
}

func (h *UserHandler) Search(res http.ResponseWriter, req *http.Request) {

	//TODO: refactor into an helper function 
	page := strings.Trim(req.URL.Query().Get("page_no"), " ")
	limit := strings.Trim(req.URL.Query().Get("limit"), " ")

	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			log.Printf("error in converting page value :%s to int\n", page)
			pageInt = 1
		}
		h.pagination.SetPage(pageInt)
	}

	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			log.Printf("error in converting limit value :%s to int\n", limit)
			limitInt = 10
		}
		h.pagination.SetLimit(limitInt)
	}

	queryStr := req.URL.Query().Get("q")
	data := make([]interface{}, 0)
	if queryStr == "" {
		utils.SendResponse(res, "failed", messages.EmailNotProvided, data, h.pagination.GetPaginationData(), http.StatusBadRequest)
		return
	}

	// e.g of query /users?q=email:johndoe@gmail.com dob:2020/11/01-2022/02/01
	// get email and dob range from query
	queries := strings.Split(queryStr, " ")
	pairs := map[string]string{}

	for _, query := range queries {
		pair := strings.Split(query, ":")
		key := strings.Trim(pair[0], " ")
		value := strings.Trim(pair[1], " ")
		pairs[key] = value
	}

	if email, ok := pairs["email"]; ok {
		// search by email
		user, err := h.userService.GetByEmail(email)
		if err != nil {
			if errors.Is(err, h.userService.ErrDataNotFound) {
				utils.SendResponse(res, "failed", messages.NoResultForSearchQuery, data, h.pagination.GetPaginationData(), http.StatusNotFound)
				return
			}
			log.Printf("an error occurred in user handler: %v", err)
			utils.SendResponse(res, "failed", messages.ErrorOccurred, data, h.pagination.GetPaginationData(), http.StatusInternalServerError)
			return
		}
		data = append(data, dto.GetUserDTOFromUser(user))
		utils.SendResponse(res, "success", "Search by email", data, h.pagination.GetPaginationData(), http.StatusOK)
		return
	}

	if dob, ok := pairs["dob"]; ok {
		// search by dob
		dateStrings := strings.Split(dob, "-")
		startDateString := strings.ReplaceAll(dateStrings[0], "/", "-")
		endDateString := strings.ReplaceAll(dateStrings[1], "/", "-")

		if startDateString == "" || endDateString == "" {
			utils.SendResponse(res, "failed", messages.ValidationError, data, h.pagination.GetPaginationData(), http.StatusBadRequest)
			return
		}

		startDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			log.Printf("error in formating the following date string: %s into a date. Due to : %v\n", startDateString, err)
			utils.SendResponse(res, "failed", messages.ValidationError, data, h.pagination.GetPaginationData(), http.StatusBadRequest)
			return
		}

		endDate, err := time.Parse("2006-01-02", endDateString)
		if err != nil {
			log.Printf("error in formating the following date string: %s into a date. Due to : %v\n", endDateString, err)
			utils.SendResponse(res, "failed", messages.ValidationError, data, h.pagination.GetPaginationData(), http.StatusBadRequest)
			return
		}
		log.Println(startDate)
		log.Println(endDate)

		// process the query
		users, err := h.userService.GetBtwDates(startDate, endDate, h.pagination.Page, h.pagination.Limit)
		if err != nil {
			log.Printf("an error occurred in user handler: %v", err)
			utils.SendResponse(res, "failed", messages.ErrorOccurred, data, h.pagination.GetPaginationData(), http.StatusInternalServerError)
			return
		}
		totalCount, err := h.userService.GetBtwDatesTotalCount(startDate, endDate, h.pagination.Page, h.pagination.Limit)
		if err != nil {
			log.Printf("an error occurred in user handler: %v", err)
			utils.SendResponse(res, "failed", messages.ErrorOccurred, data, h.pagination.GetPaginationData(), http.StatusInternalServerError)
			return
		}
		h.pagination.SetTotalCount(totalCount)
		for _, user := range users {
			data = append(data, dto.GetUserDTOFromUser(user))
		}

		utils.SendResponse(res, "success", "Search between dates", data, h.pagination.GetPaginationData(), http.StatusOK)
		return

	}
	utils.SendResponse(res, "failed", messages.SearchKeyNotProvided, data, h.pagination.GetPaginationData(), http.StatusBadRequest)

	/*
		log.Println(queryStr)
		log.Printf("%v\n", pairs)
		return
	*/
}
