package tags

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
	apis := router.Group("/tags")
	{
		apis.GET("", GetAllTags)
		apis.GET("/:tagID", GetTagByID)
		apis.PUT("/:tagID", UpdateTag)
		apis.POST("", PostTag)
		apis.DELETE("/:tagID", DeleteTag)
	}

	return apis
}

// GetAllTags godoc
func GetAllTags(c *gin.Context) {
	l.Started("GetAllTags")
	pagination := common.GetPaginationFromContext(c)
	tm := TagModel{}
	tags, err := tm.GetAllTags(pagination)
	if err != nil {
		l.Error(err.Error())
		if err == common.ErrorQuery || err == common.ErrorScanning {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// success
	c.JSON(http.StatusOK, gin.H{"tags": tags})
	l.Completed("GetAllTags")
}

// GetTagByID godoc
func GetTagByID(c *gin.Context) {
	l.Started("GetTagByID")
	ID, err := common.GetIDFromURL(c, "tagID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tm := TagModel{}
	tag, err := tm.GetTagByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No tag found with the ID: %v", ID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
	l.Completed("GetTagByID")
}

// PostTag godoc
func PostTag(c *gin.Context) {
	l.Started("PostTag")
	var tag Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tm := TagModel{}
	err := tm.PostTag(tag)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("PostTag")
}

// UpdateTag godoc
func UpdateTag(c *gin.Context) {
	l.Started("UpdateTag")
	ID, err := common.GetIDFromURL(c, "tagID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var tag Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tm := TagModel{}
	err = tm.UpdateTag(ID, tag)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := tm.GetTagByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			l.Errorf("ErrNoRows: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No tag found with the ID: %v", ID)})
			return
		}
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
	l.Completed("UpdateTag")
}

// DeleteTag godoc
func DeleteTag(c *gin.Context) {
	l.Started("DeleteTag")
	ID, err := common.GetIDFromURL(c, "tagID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tm := TagModel{}
	err = tm.DeleteTag(ID)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No tag found with the ID: %v", ID)})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	l.Completed("DeleteTag")
}
