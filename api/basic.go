package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soren/utils"
)

func GetBasic(c *gin.Context) {
	disks := utils.GetDisks()
	c.JSON(http.StatusOK, gin.H{
		"disks": disks,
	})
}
