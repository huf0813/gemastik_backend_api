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

type CartHandler struct {
	CartUseCase domain.CartUseCase
	AppDriver   *infrastructure.DriverAppService
}

func NewCartHandler(userGroup *echo.Group, c domain.CartUseCase, a *infrastructure.DriverAppService) {
	handler := &CartHandler{
		CartUseCase: c, AppDriver: a,
	}
	userGroup.POST("/cart", handler.AddProductToCart)
	userGroup.DELETE("/cart/:cart_id/product/:product_id", handler.DeleteProductFromCart)
}

func (cart *CartHandler) AddProductToCart(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	extractedToken, err := middleware.NewExtractToken(cart.AppDriver.Secret, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	create := new(domain.AddProductToCartRequest)
	if err := c.Bind(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := cart.CartUseCase.AddProductToCart(create, extractedToken.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("add product to cart", nil))
}

func (cart *CartHandler) DeleteProductFromCart(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	extractedToken, err := middleware.NewExtractToken(cart.AppDriver.Secret, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	cartID := c.Param("cart_id")
	ParseCartID, err := strconv.ParseInt(cartID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	productID := c.Param("product_id")
	ParseProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := cart.CartUseCase.DeleteProductFromCart(extractedToken.UserID, ParseProductID, ParseCartID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("delete cart successfully", nil))
}
