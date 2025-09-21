package server

import (
	"context"
	"fmt"
	"hello-go/internal/creating"
	"hello-go/internal/platform/server/handler/courses"
	"hello-go/internal/platform/server/handler/health"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddress string
	engine      *gin.Engine

	shutdownTimeout time.Duration

	creatingCourseService creating.CourseService
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, creatingCourseService creating.CourseService) (context.Context, Server) {
	server := Server{
		engine:      gin.New(),
		httpAddress: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		creatingCourseService: creatingCourseService,
	}

	server.registerRoutes()

	return serverContext(ctx), server
}

func (server *Server) Run(ctx context.Context) error {
	log.Println("Starting server on", server.httpAddress)

	srv := &http.Server{
		Addr:    server.httpAddress,
		Handler: server.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.httpAddress, err)
		}
	}()

	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), server.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutdown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()
	return ctx
}

func (server *Server) registerRoutes() {
	server.engine.GET("/health", health.CheckHandler())
	server.engine.POST("/courses", courses.CreateHandler(server.creatingCourseService))
}
