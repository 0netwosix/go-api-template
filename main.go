package main

import (
	"log"
	"github.com/0netwosix/go-api-template/api"
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
