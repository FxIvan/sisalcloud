package websocket

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func WebSocket(w http.ResponseWriter, r *http.Request){
	conn,_ := upgrader.Upgrade(	w ,	r ,	nil )

	for{
		msgType, msg , err := conn.ReadMessage()
		if err !=nil{
			return	
		}
		fmt.Printf("%s enviado: %s\n", conn.RemoteAddr(), string(msg))
		if err = conn.WriteMessage(msgType,msg); err != nil{
			return
		}
	}
}