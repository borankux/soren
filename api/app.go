package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soren/utils"
)

type App struct {
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	Version         string `json:"version"`
	Publisher       string `json:"publisher"`
	InstallLocation string `json:"install_location"`
}

func GetApps(c *gin.Context) {
	apps := utils.QueryApps()
	appList := []App{}
	for _, app := range apps {
		newApp := App{
			Name:            app.DisplayName,
			Icon:            app.DisplayIcon,
			Version:         app.DisplayVersion,
			Publisher:       app.Publisher,
			InstallLocation: app.InstallLocation,
		}
		appList = append(appList, newApp)
	}

	c.JSON(http.StatusOK, &appList)
}
