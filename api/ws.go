package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"soren/ws"
)

func WS(c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}
	ws.Manager.Register <- client
	go client.Read()
	go client.Write()
}
