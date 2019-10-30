package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIDFromURL fetches are named parameter from a Gin context and coverts it to an int.
func GetIDFromURL(c *gin.Context, paramKey string) (int, error) {
	id, err := strconv.Atoi(c.Param(paramKey))

	if err != nil {
		return -1, ErrorIntConverstion
	}

	return id, nil
}
