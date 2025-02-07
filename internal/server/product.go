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

func (s *EchoServer) UpdateProduct(ctx echo.Context) error {
	ID := ctx.Param("id")
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != product.ProductID {
		return ctx.JSON(http.StatusBadRequest, "the id in path does not equal to to the id in the body")
	}
	product, err := s.DB.UpdateProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) DeleteProduct(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteProduct(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
