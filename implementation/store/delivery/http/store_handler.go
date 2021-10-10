package http

import (
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"github.com/huf0813/gemastik_api_backend_supabase/middleware"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StoreHandler struct {
	StoreUseCase domain.StoreUseCase
	AppDriver    *infrastructure.DriverAppService
}

func NewStoreHandler(userGroup *echo.Group, s domain.StoreUseCase, a *infrastructure.DriverAppService) {
	handler := &StoreHandler{StoreUseCase: s, AppDriver: a}
	userGroup.POST("/claim_store", handler.ClaimStore)
}

func (s *StoreHandler) ClaimStore(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	token, err := middleware.NewExtractToken(s.AppDriver.Secret, authorization)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	create := new(domain.CreateStoreRequest)
	if err := c.Bind(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(create); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := s.StoreUseCase.ClaimStore(create, token.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utility.NewSuccessResponse("claim store successfully", nil))
}
