package main

import (
	"log"
	"reidev-api/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		log.Fatal()
	}
}
