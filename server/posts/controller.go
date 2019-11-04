package posts

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
	apis := router.Group("/posts")
	{
		apis.GET("", GetAllPosts)
		apis.GET("/:postID", GetPostByID)
		apis.PATCH("/:postID/postview", UpdatePostViewByID)
		apis.GET("/:postID/tags", GetPostTags)
		apis.POST("/:postID/tags", AddPostTags)
		apis.PUT("/:postID", UpdatePost)
		apis.POST("", PostPost)
		apis.DELETE("/:postID", DeletePost)
	}

	return apis
}

// GetAllPosts godoc
func GetAllPosts(c *gin.Context) {
	l.Started("GetAllPosts")
	pagination := common.GetPaginationFromContext(c)
	pm := PostModel{}
	posts, err := pm.GetAllPosts(pagination)
	if err != nil {
		l.Error(err.Error())
		if err == common.ErrorQuery || err == common.ErrorScanning {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// success
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	l.Completed("GetAllPosts")
}

// GetPostByID godoc
func GetPostByID(c *gin.Context) {
	l.Started("GetPostByID")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	post, err := pm.GetPostByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No post found with the ID: %v", ID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
	l.Completed("GetPostByID")
}

// PostPost godoc
func PostPost(c *gin.Context) {
	l.Started("PostPost")
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err := pm.PostPost(post)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("PostPost")
}

// UpdatePost godoc
func UpdatePost(c *gin.Context) {
	l.Started("UpdatePost")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.UpdatePost(ID, post)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := pm.GetPostByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			l.Errorf("ErrNoRows: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No post found with the ID: %v", ID)})
			return
		}
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
	l.Completed("UpdatePost")
}

// DeletePost godoc
func DeletePost(c *gin.Context) {
	l.Started("DeletePost")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.DeletePost(ID)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No post found with the ID: %v", ID)})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	l.Completed("DeletePost")
}

// UpdatePostViewByID godoc
func UpdatePostViewByID(c *gin.Context) {
	l.Started("UpdatePostView")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pView PostView
	if err := c.ShouldBindJSON(&pView); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.UpdatePostViewByID(ID, pView)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "updated"})
	l.Completed("UpdatePostView")
}

// GetPostTags godoc
func GetPostTags(c *gin.Context) {
	l.Started("GetPostTags")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	pt, err := pm.GetPostTags(ID)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &pt)
	l.Completed("GetPostTags")
}

// AddPostTags godoc
func AddPostTags(c *gin.Context) {
	l.Started("AddPostTags")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pt PostTags
	if err := c.ShouldBindJSON(&pt); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.AddPostTags(ID, pt)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("AddPostTags")
}
