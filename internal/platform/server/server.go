package server

import (
	"fmt"
	core "hello-go/internal"
	"hello-go/internal/platform/server/handler/courses"
	"hello-go/internal/platform/server/handler/health"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct{
	httpAddress string
	engine *gin.Engine
	
	courseRepository core.CourseRepository
}

func New(host string, port uint, courseRepository core.CourseRepository) Server {
	server := Server{
		engine: gin.New(),
		httpAddress: fmt.Sprintf("%s:%d", host, port),
		courseRepository: courseRepository,
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
	server.engine.POST("/courses", courses.CreateHandler(server.courseRepository))
}

