package router

import (
	"github.com/Wefdzen/ServMon/internal/website/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("./internal/website/static/html/*")
	r.GET("/", handlers.MainPage())
	r.GET("/test/:numServ", handlers.Graphs())
	return r
}
