package main

import (
	"github.com/gin-gonic/gin"
	"soren/api"
	"soren/middleware"
)

func main() {
	app := gin.Default()
	app.Use(middleware.Cors())
	app.GET("/info", api.GetBasic)
	app.GET("/apps", api.GetApps)
	app.GET("/ws", api.WS)
	app.GET("/test", api.Test)
	err := app.Run(":8080")
	if err != nil {
		return
	}
}

