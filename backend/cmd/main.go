package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/websocket"
	"github.com/LakeBlair/royalgame/backend/server"
)



func main() {
	fmt.Println("Starting the server")

	config := server.Get("config.yaml")
	CORS := config.GetStringSlice("CORS")
	upgrader := websocket.Upgrader{}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		check := server.ContainsString(CORS, r.URL.Host)
		if !check {
			log.Printf("Error Origin: Host - (%s) | URL - (%s)", r.URL.Host, r.URL)
		}
		return true
	}

	wsHandler := server.NewWSHandler(upgrader)
	log.Println(config.GetString("ADDR"))
	wsHandler.LaunchGame()
}