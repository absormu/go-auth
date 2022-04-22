package pagination

import (
	"strconv"

	"github.com/absormu/go-auth/app/entity"
	md "github.com/absormu/go-auth/app/middleware"
	cm "github.com/absormu/go-auth/pkg/configuration"
	"github.com/labstack/echo/v4"
)

type GeneratePagination struct {
	Total       *int64 `json:"total"`
	PerPage     int64  `json:"per_page"`
	CurrentPage int64  `json:"current_page"`
	LastPage    *int64 `json:"last_page"`
}

func Pagination(c echo.Context) (meta entity.Meta, e error) {
	logger := md.GetLogger(c)
	logger.Info("pkg: Pagination")

	// Get parameter limit
	limitStr := c.QueryParam("limit")
	var limit int64
	if limitStr != "" {
		limit, e = strconv.ParseInt(limitStr, 10, 64)
		if e == nil {
			if (limit == 0) || (limit > cm.Config.LimitQuery) {
				limit = cm.Config.LimitQuery
			}
		} else {
			logger.WithField("error", e.Error()).Error("Catch error failure QueryParams")
			return
		}
	} else {
		limit = cm.Config.LimitQuery
	}
	meta.Limit = limit

	// Get parameter page
	pageStr := c.QueryParam("page")
	var page int64
	if pageStr != "" {
		page, e = strconv.ParseInt(pageStr, 10, 64)
		if e == nil {
			if page == 0 {
				page = 1
			}
		} else {
			logger.WithField("error", e.Error()).Error("Catch error failure QueryParams")
			return
		}
	} else {
		page = 1
	}
	meta.Page = page

	if page > 1 {
		meta.Offset = limit * (page - 1)
	}

	// Get parameter pagination
	pagination := false
	paginationStr := c.QueryParam("pagination")
	if paginationStr != "" {
		pagination, e = strconv.ParseBool(paginationStr)
		if e != nil {
			logger.WithField("error", e.Error()).Error("Catch error failure QueryParams")
			return
		}
	}
	meta.Pagination = pagination

	return
}

func GenerateMeta(c echo.Context, total int64, limit int64, page int64, offset int64, pagination bool, params map[string]string) GeneratePagination {
	var meta GeneratePagination
	queryParam := ""

	if len(params) > 0 {
		i := 0
		queryParam += "&"
		for index, query := range params {
			queryParam += index + "=" + query
			if (len(params) - 1) > i {
				queryParam += "&"
			}
			i++
		}
	}

	if pagination == true {
		meta.Total = &total
		meta.PerPage = limit
		lastPage := total / limit
		meta.LastPage = &lastPage
		if (total % limit) > 0 {
			*meta.LastPage++
		}
		if page == 0 {
			meta.CurrentPage = 1
		} else {
			meta.CurrentPage = page
		}
	}

	return meta
}
