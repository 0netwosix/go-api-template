package api

import (
	"github.com/gin-gonic/gin"
)

// Server serves http requests for the project
type Server struct {
	router *gin.Engine
}

// NewServer creates a new http server and setup routing
func NewServer() (*Server, error) {
	var server = new(Server)

	router := gin.Default()

	router.GET("/", server.index)

	server.router = router

	return server, nil
}

// Start runs the http server on the given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
