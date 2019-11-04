package main

import (
	"net/http"

	"github.com/codingmechanics/applogger"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/masihur1989/masihurs-blog/server/categories"
	"github.com/masihur1989/masihurs-blog/server/common"
	"github.com/masihur1989/masihurs-blog/server/posts"
	"github.com/masihur1989/masihurs-blog/server/tags"
	"github.com/masihur1989/masihurs-blog/server/users"
)

var l applogger.Logger

func setupRouter() *gin.Engine {
	// new gin engine
	// custom gin engine
	router := gin.New()
	router.Use(l.GinLogger())
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	// setup route group
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		users.RegisterRoutes(v1)
		categories.RegisterRoutes(v1)
		tags.RegisterRoutes(v1)
		posts.RegisterRoutes(v1)
	}
	// return the router
	return router
}

func main() {
	l.DisableColor = true
	// start logging
	l.Start(applogger.LevelDebug)
	// load configs
	common.ConfigureApp()

	r := setupRouter()
	// start and run the server
	_ = r.Run(":3000")
	// close db connection
	defer common.CloseDB()
	// stop logging
	l.Stop()
}
