package main

import (
	"log"
	"github.com/gorilla/websocket"
)

func main() {
	// Configure the dialer
	dialer := websocket.Dialer{}

	// Connect to the WebSocket server
	conn, _, err := dialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	// Send a message to the server
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hi From the Client!"))
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	// Continuously read messages from the server
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received from server: %s\n", message)
	}

	// Ensure the connection is closed on exit
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing WebSocket connection:", err)
		}
	}()
}

