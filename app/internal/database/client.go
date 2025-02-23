package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetCustomerById(ctx context.Context, ID string) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, ID string) error

	GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, ID string) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, ID string) error

	GetAllServices(ctx context.Context) ([]models.Service, error)
	AddService(ctx context.Context, service *models.Service) (*models.Service, error)
	GetServiceById(ctx context.Context, ID string) (*models.Service, error)
	UpdateService(ctx context.Context, service *models.Service) (*models.Service, error)
	DeleteService(ctx context.Context, ID string) error

	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
	AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	GetVendorById(ctx context.Context, ID string) (*models.Vendor, error)
	UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	DeleteVendor(ctx context.Context, ID string) error
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbTablePrefix := os.Getenv("DB_TABLE_PREFIX")
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" || dbTablePrefix == "" {
		return nil, errors.New("one or more environment variables related to DB has not been set")
	}
	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, errors.New("invalid db port")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPortInt, dbUser, dbPass, dbName, "disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbTablePrefix + ".",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
