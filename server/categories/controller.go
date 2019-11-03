package categories

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/codingmechanics/applogger"
	"github.com/gin-gonic/gin"
	"github.com/masihur1989/masihurs-blog/server/common"
)

var l applogger.Logger

// RegisterRoutes - Register all the routes for users
func RegisterRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	apis := router.Group("/categories")
	{
		apis.GET("", GetAllCategories)
		apis.GET("/:categoryID", GetCategoryByID)
		apis.PUT("/:categoryID", UpdateCategory)
		apis.POST("", PostCategory)
		apis.DELETE("/:categoryID", DeleteCategory)
	}

	return apis
}

// GetAllCategories godoc
func GetAllCategories(c *gin.Context) {
	l.Started("GetAllCategories")
	pagination := common.GetPaginationFromContext(c)
	cm := CategoryModel{}
	categories, err := cm.GetAllCategories(pagination)
	if err != nil {
		l.Error(err.Error())
		if err == common.ErrorQuery || err == common.ErrorScanning {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// success
	c.JSON(http.StatusOK, gin.H{"categories": categories})
	l.Completed("GetAllCategories")
}

// GetCategoryByID godoc
func GetCategoryByID(c *gin.Context) {
	l.Started("GetCategoryByID")
	ID, err := common.GetIDFromURL(c, "categoryID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cm := CategoryModel{}
	category, err := cm.GetCategoryByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No category found with the ID: %v", ID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
	l.Completed("GetCategoryByID")
}

// PostCategory godoc
func PostCategory(c *gin.Context) {
	l.Started("PostCategory")
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cm := CategoryModel{}
	err := cm.PostCategory(category)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("PostCategory")
}

// UpdateCategory godoc
func UpdateCategory(c *gin.Context) {
	l.Started("UpdateCategory")
	ID, err := common.GetIDFromURL(c, "categoryID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cm := CategoryModel{}
	err = cm.UpdateCategory(ID, category)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := cm.GetCategoryByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			l.Errorf("ErrNoRows: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No category found with the ID: %v", ID)})
			return
		}
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
	l.Completed("UpdateCategory")
}

// DeleteCategory godoc
func DeleteCategory(c *gin.Context) {
	l.Started("DeleteCategory")
	ID, err := common.GetIDFromURL(c, "categoryID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cm := CategoryModel{}
	err = cm.DeleteCategory(ID)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No category found with the ID: %v", ID)})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	l.Completed("DeleteCategory")
}
