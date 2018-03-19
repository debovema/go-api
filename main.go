package main

import (
	"github.com/debovema/go-api/handlers/content"
	"github.com/debovema/go-api/handlers/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	enableLog()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	accounts := getAccounts()

	basicallyAuthorized := r.Group("/basic")
	basicallyAuthorized.Use(middleware.AuthorizationHandler(middleware.Basic, basicallyAuthorized, accounts))
	{
		getRoutes(basicallyAuthorized)
	}

	jwtAuthorized := r.Group("/jwt")
	jwtAuthorized.Use(middleware.AuthorizationHandler(middleware.JWT, jwtAuthorized, accounts))
	{
		getRoutes(jwtAuthorized)
	}

	// basic ping request displaying pong in a JSON result
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}

func getAccounts() gin.Accounts {
	return map[string]string{
		"foo": "bar",
	}
}

func getRoutes(routerGroup *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		routerGroup.GET("/articles", content.ArticlesHandler()),
		routerGroup.GET("/articles/:id", content.ArticlesHandler()),
	}
}

func enableLog() {
	// gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
