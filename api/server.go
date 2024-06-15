package api

import (
	db "github.com/danielHieu/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server struct    servers HTTP request for our banking service.
type Server struct {
	store	*db.Store	
	router *gin.Engine
}

// NewServer function    create a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	// add router
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Start method    run the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) *gin.H {
	return &gin.H{"error": err.Error()}
}
