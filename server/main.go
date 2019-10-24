package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/masihur1989/masihurs-blog/server/common"
)

func setupRouter() *gin.Engine {
	// set the router with default one that comes with gin
	router := gin.Default()
	// server the static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	// setup route group
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	// return the router
	return router
}

func main() {
	common.ConfigureApp()
	r := setupRouter()
	// start and run the servier
	_ = r.Run(":3000")
	defer common.CloseDB()
}
