package main

import (
	"log"
	"lw/edr/api"
)

// TODO: add these to envs
const address string = "localhost:8080"

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
