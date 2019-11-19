package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codingmechanics/applogger"
	"github.com/dgrijalva/jwt-go"
	jwtMiddleware "github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/masihur1989/masihurs-blog/server/common"
	"golang.org/x/crypto/bcrypt"
)

var l applogger.Logger

// RegisterRoutes - Register all the routes for users
func RegisterRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	apiUsers := router.Group("/users")
	{
		apiUsers.GET("", jwtMiddleware.Auth(os.Getenv("JWT_SECRET")), GetAllUsers)
		apiUsers.GET("/:userID", GetUserByID)
		apiUsers.PATCH("/:userID/forgotpassword", ForgotPasswordController)
		apiUsers.POST("", Register)
		apiUsers.POST("/login", Login)
		apiUsers.DELETE("/:userID", DeleteUser)
	}
	return apiUsers
}

// GetAllUsers godoc
func GetAllUsers(c *gin.Context) {
	l.Started("GetAllUsers")
	pagination := common.GetPaginationFromContext(c)
	um := UserModel{}
	users, err := um.GetAllUsers(pagination)
	if err != nil {
		l.Errorf("GetAllUsers: ", err)
		if err == common.ErrorQuery || err == common.ErrorScanning {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// success
	c.JSON(http.StatusOK, gin.H{"users": users})
	l.Completed("GetAllUsers")
}

// GetUserByID godoc
func GetUserByID(c *gin.Context) {
	l.Started("GetUserByID")
	ID, err := common.GetIDFromURL(c, "userID")
	if err != nil {
		l.Error(common.ErrorIntConverstion.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	user, err := um.GetUserByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No user found with the ID: %v", ID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	l.Completed("GetUserByID")
}

// Login godoc
func Login(c *gin.Context) {
	l.Started("Login")
	var ul UserLogin
	if err := c.ShouldBindJSON(&ul); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	user, err := um.PostLogin(ul)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case err == common.ErrorPasswordMatching:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		default:
			l.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	claims := common.Claims{
		user.Name,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	}
	t, err := common.GeneratToken(claims)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": t})
	l.Completed("Login")
}

// Register godoc
func Register(c *gin.Context) {
	l.Started("Register")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	err := um.PostUser(user)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
	l.Completed("Register")
}

// ForgotPasswordController godoc
func ForgotPasswordController(c *gin.Context) {
	l.Started("ForgotPasswordController")
	ID, err := common.GetIDFromURL(c, "userID")
	if err != nil {
		l.Errorf("ERROR PARSING USERID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payload ForgotPassword

	if err := c.ShouldBindJSON(&payload); err != nil {
		l.Errorf("ERROR BINDING PAYLOAD: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Password == payload.NewPassword {
		l.Error("No Change in Password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot use the same password for new password!"})
		return
	}

	um := UserModel{}
	err = um.UpdateUserPassword(ID, payload)

	if err == sql.ErrNoRows {
		l.Errorf("ErrNoRows: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No user found with the userID: %v", ID)})
		return
	} else if err == bcrypt.ErrHashTooShort || err == bcrypt.ErrMismatchedHashAndPassword {
		l.Errorf("Bcrypt Hashing Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Wrong Password")})
		return
	}

	response, err := um.GetUserByID(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			l.Errorf("ErrNoRows: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No user found with the userID: %v", ID)})
			return
		}
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	l.Completed("ForgotPasswordController")
}

// DeleteUser godoc
func DeleteUser(c *gin.Context) {
	l.Started("DeleteUser")
	ID, err := common.GetIDFromURL(c, "userID")
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	um := UserModel{}
	err = um.DeleteUser(ID)
	if err != nil && err == sql.ErrNoRows {
		l.Errorf("ErrNoRows: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	l.Completed("DeleteUser")
}
