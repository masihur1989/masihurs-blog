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
		apis.PUT("/:postID", UpdatePost)
		apis.POST("", PostPost)
		apis.DELETE("/:postID", DeletePost)
		// postview endpoint
		apis.PATCH("/:postID/postview", UpdatePostViewByID)
		// tags endpoint
		apis.GET("/:postID/tags", GetPostTags)
		apis.POST("/:postID/tags", AddPostTags)
		// like endpoint
		apis.GET("/:postID/likes", GetPostLikes)
		apis.POST("/:postID/likes", LikePost)
		// comment endpoint
		apis.GET("/:postID/comments", GetAllPostComments)
		apis.POST("/:postID/comments", CommentPost)
		apis.PUT("/:postID/comments/:commentID", UpdateComment)
		apis.GET("/:postID/comments/:commentID", GetPostComment)
		apis.DELETE("/:postID/comments/:commentID", DeleteComment)
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

// GetPostLikes godoc
func GetPostLikes(c *gin.Context) {
	l.Started("PostLikes")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	pt, err := pm.GetPostLikes(ID)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &pt)
	l.Completed("PostLikes")
}

// LikePost godoc
func LikePost(c *gin.Context) {
	l.Started("LikePost")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pl PostLike
	if err := c.ShouldBindJSON(&pl); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	if pl.Like {
		err = pm.LikePost(ID, pl)
		if err != nil {
			l.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		err = pm.DeletePostLike(ID, pl)
		if err != nil {
			l.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("LikePost")
}

// GetAllPostComments godoc
func GetAllPostComments(c *gin.Context) {
	l.Started("GetAllPostComments")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	pc, err := pm.GetPostComments(ID)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post_comments": pc})
	l.Completed("GetAllPostComments")
}

// CommentPost godoc
func CommentPost(c *gin.Context) {
	l.Started("CommentPost")
	ID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pc PostComment
	if err := c.ShouldBindJSON(&pc); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.AddComment(ID, pc)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("CommentPost")
}

// UpdateComment godoc
func UpdateComment(c *gin.Context) {
	l.Started("UpdateComment")
	postID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentID, err := common.GetIDFromURL(c, "commentID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pc PostComment
	if err := c.ShouldBindJSON(&pc); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.UpdateComment(postID, commentID, pc)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("UpdateComment")
}

// GetPostComment godoc
func GetPostComment(c *gin.Context) {
	l.Started("GetPostComment")
	postID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentID, err := common.GetIDFromURL(c, "commentID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pm PostModel
	pc, err := pm.GetPostComment(postID, commentID)

	c.JSON(http.StatusOK, &pc)
	l.Completed("GetPostComment")
}

// DeleteComment godoc
func DeleteComment(c *gin.Context) {
	l.Started("DeleteComment")
	postID, err := common.GetIDFromURL(c, "postID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentID, err := common.GetIDFromURL(c, "commentID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pm := PostModel{}
	err = pm.DeleteComment(postID, commentID)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No comment found with IDs.")})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	l.Completed("DeleteComment")
}
