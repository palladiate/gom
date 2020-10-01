package main

import (
	"github.com/palladiate/gom/player"
	"log"
	"net"
)

type Server struct {
	Host string
	Port string
	Type string
}

func main() {
	gom := Server{
		Host: "localhost",
		Port: "9000",
		Type: "tcp",
	}

	listener, err := net.Listen(gom.Type, gom.Host+":"+gom.Port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s, port %s", gom.Host, gom.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		toon := player.NewPlayer(conn)
		go toon.Play()
	}
}
