package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/masihur1989/masihurs-blog/server/common"
)

// RegisterRoutes - Register all the routes for users
func RegisterRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	apiUsers := router.Group("/users")
	{
		apiUsers.GET("", GetAllUsers)
		apiUsers.GET("/:userID", GetUserByID)
		apiUsers.PATCH("/:userID/forgotpassword", ForgotPasswordController)
		apiUsers.POST("", Register)
		apiUsers.DELETE("/:userID", DeleteUser)
	}

	return apiUsers
}

// GetAllUsers godoc
func GetAllUsers(c *gin.Context) {
	pagination := common.GetPaginationFromContext(c)
	um := UserModel{}
	users, err := um.GetAllUsers(pagination)
	if err != nil {
		log.Printf("GetAllUsers: %v\n", err)
		if err == common.ErrorQuery || err == common.ErrorScanning {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// success
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserByID godoc
func GetUserByID(c *gin.Context) {
	ID, err := common.GetIDFromURL(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	user, err := um.GetUserByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No user found with the userID: %v", ID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Register godoc
func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.Error(err)
		log.Printf("JSON BINDING ERROR: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	err := um.PostUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

// ForgotPasswordController -
func ForgotPasswordController(c *gin.Context) {
	ID, err := common.GetIDFromURL(c, "userID")
	if err != nil {
		log.Printf("ERROR PARSING USERID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payload ForgotPassword

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Printf("ERROR BINDING PAYLOAD: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Password == payload.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot use the same password for new password!"})
		return
	}

	um := UserModel{}
	err = um.UpdateUserPassword(ID, payload)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No user found with the userID: %v", ID)})
		return
	} else if err == bcrypt.ErrHashTooShort || err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Password")})
		return
	}

	response, err := um.GetUserByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No user found with the userID: %v", ID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser -
func DeleteUser(c *gin.Context) {
	ID, err := common.GetIDFromURL(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	err = um.DeleteUser(ID)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
