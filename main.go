package main

import (
	"zmos-underground-monitor/conf"
	"zmos-underground-monitor/pkg/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.GET("/server/menu", api.GetServerMenu)
	v1.GET("/server/info", api.GetServerInfo)
	router.Run(":" + conf.Conf.ApiServer.Port)
}
