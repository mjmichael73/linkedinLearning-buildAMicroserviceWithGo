package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/database"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
	GetAllCustomers(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK", Message: "Server is ready."})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure", Message: "Server is not ready."})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK", Message: "Server is live"})
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)
}
