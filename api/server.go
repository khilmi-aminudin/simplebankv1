package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/khilmi-aminudin/simplebankv1/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}

	// Register The Currency Validation on binding Tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
