package http

import (
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"github.com/huf0813/gemastik_api_backend_supabase/middleware"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceUseCase domain.InvoiceUseCase
	AppDriver      *infrastructure.DriverAppService
}

func NewInvoiceHandler(userGroup *echo.Group, i domain.InvoiceUseCase, a *infrastructure.DriverAppService) {
	handler := &InvoiceHandler{InvoiceUseCase: i, AppDriver: a}
	userGroup.POST("/invoice/create", handler.CreateInvoice)
	userGroup.PATCH("/invoice/:code/shipped", handler.UpdateInvoiceShipped)
	userGroup.PATCH("/invoice/:code/accepted", handler.UpdateInvoiceAccepted)
	userGroup.POST("/invoice/product_review", handler.CreateInvoiceProductReview)
}

func (i *InvoiceHandler) CreateInvoice(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	token, err := middleware.NewExtractToken(i.AppDriver.Secret, authorization)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	create := new(domain.CreateInvoiceRequest)
	if err := c.Bind(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	code, err := i.InvoiceUseCase.CreateInvoice(create, token.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("invoice created successfully", map[string]string{
		"invoice_code": code,
	}))
}

func (i *InvoiceHandler) UpdateInvoiceShipped(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	extractedToken, err := middleware.NewExtractToken(i.AppDriver.Secret, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	code := c.Param("code")
	if err := i.InvoiceUseCase.UpdateInvoiceStatus(2, extractedToken.UserID, code); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("updated invoice to shipped", nil))
}

func (i *InvoiceHandler) UpdateInvoiceAccepted(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	extractedToken, err := middleware.NewExtractToken(i.AppDriver.Secret, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	code := c.Param("code")
	if err := i.InvoiceUseCase.UpdateInvoiceStatus(3, extractedToken.UserID, code); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("updated invoice to accepted", nil))
}

func (i *InvoiceHandler) CreateInvoiceProductReview(c echo.Context) error {
	create := new(domain.CreateInvoiceProductReviewRequest)
	if err := c.Bind(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := i.InvoiceUseCase.CreateInvoiceProductReview(create); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("create product review successfully", nil))
}
