package handlers

import (
	"fmt"
	"net/http"

	"github.com/Wefdzen/ServMon/internal/website/services"
	"github.com/Wefdzen/ServMon/pkg/service"
	"github.com/gin-gonic/gin"
)

func MainPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	}
}

func GraphsWithMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.DrawAllParamWithMode(c)
	}
}

type ServersName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetServersNames() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, err := service.GetInfoServers("./servers.json")
		if err != nil {
			fmt.Println(err)
		}
		var res []ServersName
		count := 1 //localhost:port/test/count/mode
		for i := 0; i < len(tmp); i++ {
			res = append(res, ServersName{Id: count, Name: tmp[i].NameOfService})
			count++
		}

		c.JSON(http.StatusOK, res)
	}
}
