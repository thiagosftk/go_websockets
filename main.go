package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	println("Testando esse tal do WebSocket")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}
	for {
		// Vamos ler a mensagem recebida via Websocket
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Logando no console do Webserver
		fmt.Println("Mensagem recebida: ", string(msg))

		// Devolvendo a mensagem recebida de volta para o cliente
		err = socket.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
