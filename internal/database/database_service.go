package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/dberrors"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).Find(&services)
	return services, result.Error
}

func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
	service.ServiceID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return service, nil
}

func (c Client) GetServiceById(ctx context.Context, ID string) (*models.Service, error) {
	service := &models.Service{}
	result := c.DB.WithContext(ctx).Where(&models.Service{ServiceID: ID}).First(&service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "service", ID: ID}
		}
		return nil, result.Error
	}
	return service, nil
}
