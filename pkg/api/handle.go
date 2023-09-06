package api

import (
	"net/http"
	"zmos-underground-monitor/pkg/service"

	"github.com/gin-gonic/gin"
)

func GetServerMenu(c *gin.Context) {
	sm := service.NewServerMenu()
	sm.SetData()
	c.JSON(http.StatusOK, sm)
}

func GetServerInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "reboot sucess")
}
