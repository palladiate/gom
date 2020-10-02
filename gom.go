package main

import (
	"github.com/palladiate/gom/listeners"
	"log"
)

type Server struct {
	Host string
	Port string
	Type string
}

func main() {
	gom := listeners.NewTelnet("localhost", 9000)

	go gom.Start()
	defer gom.Stop()
	var gomPlayers = gom.PlayerChannel()

	log.Printf("Server listening on %s, port %v", gom.Host, gom.Port)

	for {
		select {
		case toon := <- gomPlayers:
			log.Print("Player returned.")
			go toon.Play()
		}
	}
}
