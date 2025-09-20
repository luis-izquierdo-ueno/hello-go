package server

import (
	"fmt"
	"hello-go/internal/platform/server/handler/health"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct{
	httpAddress string
	engine *gin.Engine
}

func New(host string, port uint) Server {
	server := Server{
		engine: gin.New(),
		httpAddress: fmt.Sprintf("%s:%d", host, port),
	}

	server.registerRoutes()

	return server
}

func (server *Server) Run() error {
	log.Println("Starting server on", server.httpAddress)
	return server.engine.Run(server.httpAddress)
}

func (server *Server) registerRoutes() {
	server.engine.GET("/health", health.CheckHandler())
}

