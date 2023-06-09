package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kombolewis/kompani/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/company", server.createCompany)
	router.GET("/company/:id", server.getCompany)
	router.GET("/company", server.listCompanies)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start() error {
	return server.router.Run()
}
