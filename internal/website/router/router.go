package router

import (
	"github.com/Wefdzen/ServMon/internal/website/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//clear log req
	ginMode := "release"
	gin.SetMode(ginMode)
	r := gin.New() //if you want gin http log in console set gin.Default and delete what 1 row below and 2 above
	r.Use(gin.Recovery())

	r.Static("/static", "./internal/website/static/html")
	r.LoadHTMLGlob("./internal/website/static/html/*")

	r.GET("/", handlers.MainPage())
	r.GET("/test/:numServ/:mode", handlers.GraphsWithMode())
	r.GET("/api/servers", handlers.GetServersNames())
	return r
}
