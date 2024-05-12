package main

import (
	"fmt"
	"github.com/LakeBlair/royalgame/backend/internal"
	"github.com/LakeBlair/royalgame/backend/server"
)

func main() {
	fmt.Println("Starting the game")
	go server.LaunchGameServer()
	internal.Init_Game()
}