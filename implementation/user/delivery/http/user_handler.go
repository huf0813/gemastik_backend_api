package http

import (
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"github.com/huf0813/gemastik_api_backend_supabase/middleware"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
	AppDriver   *infrastructure.DriverAppService
}

func NewUserHandler(e *echo.Echo, userGroup *echo.Group, u domain.UserUseCase, a *infrastructure.DriverAppService) {
	handler := &UserHandler{UserUseCase: u, AppDriver: a}
	e.POST("/sign_up", handler.SignUp)
	e.POST("/sign_in", handler.SignIn)

	userGroup.GET("/profile", handler.GetProfile)
	userGroup.POST("/claim_supplier", handler.ClaimSupplier)
}

func (u *UserHandler) SignUp(c echo.Context) error {
	signUp := new(domain.UserSignUp)
	if err := c.Bind(signUp); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(signUp); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(signUp); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := u.UserUseCase.SignUp(signUp); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utility.NewSuccessResponse("user registered successfully", nil))
}

func (u *UserHandler) SignIn(c echo.Context) error {
	signIn := new(domain.UserSignIn)
	if err := c.Bind(signIn); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(signIn); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	token, err := u.UserUseCase.SignIn(signIn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utility.NewSuccessResponse("user signed in successfully", map[string]string{"access_token": token}))
}

func (u *UserHandler) ClaimSupplier(c echo.Context) error {
	claimSupplier := new(domain.UserClaimSupplier)
	if err := c.Bind(claimSupplier); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	if err := c.Validate(claimSupplier); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	token := c.Request().Header.Get("Authorization")
	extractedToken, err := middleware.NewExtractToken(u.AppDriver.Secret, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}

	if err := u.UserUseCase.ClaimSupplier(claimSupplier, extractedToken.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utility.NewSuccessResponse("claim supplier successfully", nil))
}

func (u *UserHandler) GetProfile(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	extractedToken, err := middleware.NewExtractToken(u.AppDriver.Secret, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewFailResponse(err.Error()))
	}
	selectedColumns := c.QueryParam("select")
	result, err := u.UserUseCase.GetProfile(extractedToken.UserID, selectedColumns)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utility.NewSuccessResponse("get profile successfully", result))
}
