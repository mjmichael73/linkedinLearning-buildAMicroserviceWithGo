package server

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/database"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/models"
)

type Server interface {
	Start() error

	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	GetCustomerById(ctx echo.Context) error
	UpdateCustomer(ctx echo.Context) error
	DeleteCustomer(ctx echo.Context) error

	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error
	DeleteProduct(ctx echo.Context) error

	GetAllServices(ctx echo.Context) error
	AddService(ctx echo.Context) error
	GetServiceById(ctx echo.Context) error
	UpdateService(ctx echo.Context) error
	DeleteService(ctx echo.Context) error

	GetAllVendors(ctx echo.Context) error
	AddVendor(ctx echo.Context) error
	GetVendorById(ctx echo.Context) error
	UpdateVendor(ctx echo.Context) error
	DeleteVendor(ctx echo.Context) error
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
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set in the env")
		return errors.New("APP_PORT is not set in the env")
	}
	if err := s.echo.Start(":" + appPort); err != nil && err != http.ErrServerClosed {
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
	cg.POST("", s.AddCustomer)
	cg.GET("/:id", s.GetCustomerById)
	cg.PUT("/:id", s.UpdateCustomer)
	cg.DELETE("/:id", s.DeleteCustomer)

	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)
	pg.POST("", s.AddProduct)
	pg.GET("/:id", s.GetProductById)
	pg.PUT("/:id", s.UpdateProduct)
	pg.DELETE("/:id", s.DeleteProduct)

	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)
	sg.POST("", s.AddService)
	sg.GET("/:id", s.GetServiceById)
	sg.PUT("/:id", s.UpdateService)
	sg.DELETE("/:id", s.DeleteService)

	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)
	vg.POST("", s.AddVendor)
	vg.GET("/:id", s.GetVendorById)
	vg.PUT("/:id", s.UpdateVendor)
	vg.DELETE("/:id", s.DeleteVendor)
}
