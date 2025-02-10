package handlers

import (
	"net/http"

	"github.com/Wefdzen/ServMon/internal/website/services"
	"github.com/gin-gonic/gin"
)

func MainPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	}
}

func Graphs() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.DrawAllParam(c)
	}
}
