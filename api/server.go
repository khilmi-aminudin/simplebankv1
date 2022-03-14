package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/khilmi-aminudin/simplebankv1/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
