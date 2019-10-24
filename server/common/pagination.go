package common

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Pagination struct
type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

const (
	defaultPage     = 1
	defaultPageSize = 25
)

// GetPaginationFromContext Scan common pagination variables from the query parameters into a struct.
func GetPaginationFromContext(c *gin.Context) Pagination {
	pagination := Pagination{
		Page:     defaultPage,
		PageSize: defaultPageSize,
	}

	err := c.Bind(&pagination)

	if err != nil {
		return pagination
	}

	if pagination.Page < defaultPage {
		pagination.Page = defaultPage
	}

	if pagination.PageSize <= 0 {
		pagination.PageSize = defaultPageSize
	}

	return pagination
}

// ApplyPaginationToQuery Apply generic pagination to a query.
// This method is meant to be used in conjunction with sqlx.Named.
func ApplyPaginationToQuery(query string, args map[string]interface{}, pagination Pagination) (string, map[string]interface{}) {
	args["offset"] = (pagination.Page - 1) * pagination.PageSize
	args["page_size"] = pagination.PageSize
	query = fmt.Sprintf(`
		%s
		LIMIT :offset, :page_size
	`, query)
	log.Printf("ApplyPaginationToQuery - QUERY: %v", query)
	return query, args
}
