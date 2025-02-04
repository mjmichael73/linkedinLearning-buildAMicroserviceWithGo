package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/dberrors"
	"github.com/mjmichael73/linkedinLearning-buildAMicroserviceWithGo/internal/models"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")
	products, err := s.DB.GetAllProducts(ctx.Request().Context(), vendorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)

}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (s *EchoServer) GetProductById(ctx echo.Context) error {
	id := ctx.Param("id")
	product, err := s.DB.GetProductById(ctx.Request().Context(), id)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}
