package router

import (
	"github.com/Wefdzen/ServMon/internal/website/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./internal/website/static/html")
	r.LoadHTMLGlob("./internal/website/static/html/*")

	r.GET("/", handlers.MainPage())
	r.GET("/test/:numServ/:mode", handlers.GraphsWithMode())
	return r
}
