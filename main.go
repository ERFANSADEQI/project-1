package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
    if err!= nil {
       http.NotFound(w, r)
	   return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err!= nil {
            break
        }

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	
    router.GET("/ws/echo", func(c *gin.Context) {
        wshandler(c.Writer, c.Request)
    })

	router.SetTrustedProxies([]string{"127.0.0.1"})
    router.Run(":8889")
}