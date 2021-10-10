package http

import (
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"github.com/huf0813/gemastik_api_backend_supabase/middleware"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductUseCase domain.ProductUseCase
	AppDriver      *infrastructure.DriverAppService
}

func NewProductHandler(userGroup *echo.Group, p domain.ProductUseCase, a *infrastructure.DriverAppService) {
	handler := &ProductHandler{ProductUseCase: p, AppDriver: a}
	userGroup.POST("/product/create", handler.CreateProduct)
	userGroup.PATCH("/product/:product_id", handler.UpdateProduct)
	userGroup.DELETE("/product/:product_id", handler.DeleteProduct)
}

func (p *ProductHandler) CreateProduct(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	token, err := middleware.NewExtractToken(p.AppDriver.Secret, authorization)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	create := new(domain.CreateProductRequest)
	if err := c.Bind(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	//return c.JSON(http.StatusOK, utility.NewSuccessResponse("created product successfully", create))

	if err := p.ProductUseCase.CreateProduct(create, token.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utility.NewSuccessResponse("created product successfully", nil))

}

func (p *ProductHandler) UpdateProduct(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	token, err := middleware.NewExtractToken(p.AppDriver.Secret, authorization)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	update := new(domain.CreateProductRequest)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(update); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	productID := c.Param("product_id")
	ParseProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := p.ProductUseCase.UpdateProduct(update, ParseProductID, token.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("updated product successfully", nil))
}

func (p *ProductHandler) DeleteProduct(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	token, err := middleware.NewExtractToken(p.AppDriver.Secret, authorization)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	productID := c.Param("product_id")
	ParseProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := p.ProductUseCase.DeleteProduct(ParseProductID, token.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("deleted product successfully", nil))
}
