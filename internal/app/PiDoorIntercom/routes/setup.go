package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./webpage", true)))
	r.Static("/css", "css")
	r.Static("/js", "js")

	r.GET("/video", Video)

	return r
}
