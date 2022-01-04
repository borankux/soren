package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soren/ws"
)

func Test(c *gin.Context) {
	ws.Manager.Send([]byte("Hello World Fucker"), nil)
	clients := ws.Manager.Clients
	var clientList []string
	for client := range clients {
		clientList = append(clientList, client.ID)
	}

	c.JSON(http.StatusOK, &clientList)
}
