package main

import (
	"app4/server"
	"log"
)

func main() {
	conn := server.ServerConnection{}
	if err := conn.NewServerConnection(); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
