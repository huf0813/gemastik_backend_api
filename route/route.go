package route

import (
	"github.com/go-playground/validator"
	_cartHandler "github.com/huf0813/gemastik_api_backend_supabase/implementation/cart/delivery/http"
	_cartRepositorySupabase "github.com/huf0813/gemastik_api_backend_supabase/implementation/cart/repository/supabase"
	_cartUseCase "github.com/huf0813/gemastik_api_backend_supabase/implementation/cart/usecase"
	_fetchHandler "github.com/huf0813/gemastik_api_backend_supabase/implementation/fetch/delivery/http"
	_fetchRepositorySupabase "github.com/huf0813/gemastik_api_backend_supabase/implementation/fetch/repository/supabase"
	_fetchUseCase "github.com/huf0813/gemastik_api_backend_supabase/implementation/fetch/usecase"
	_invoiceHandler "github.com/huf0813/gemastik_api_backend_supabase/implementation/invoice/delivery/http"
	_invoiceRepositorySupabase "github.com/huf0813/gemastik_api_backend_supabase/implementation/invoice/repository/supabase"
	_invoiceUseCase "github.com/huf0813/gemastik_api_backend_supabase/implementation/invoice/usecase"
	_productHandler "github.com/huf0813/gemastik_api_backend_supabase/implementation/product/delivery/http"
	_productRepositorySupabase "github.com/huf0813/gemastik_api_backend_supabase/implementation/product/repository/supabase"
	_productUseCase "github.com/huf0813/gemastik_api_backend_supabase/implementation/product/usecase"
	_storeHandler "github.com/huf0813/gemastik_api_backend_supabase/implementation/store/delivery/http"
	_storeRepositorySupabase "github.com/huf0813/gemastik_api_backend_supabase/implementation/store/repository/supabase"
	_storeUseCase "github.com/huf0813/gemastik_api_backend_supabase/implementation/store/usecase"
	_userHandler "github.com/huf0813/gemastik_api_backend_supabase/implementation/user/delivery/http"
	_userRepositorySupabase "github.com/huf0813/gemastik_api_backend_supabase/implementation/user/repository/supabase"
	_userUseCase "github.com/huf0813/gemastik_api_backend_supabase/implementation/user/usecase"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewRoutes(e *echo.Echo, authMiddleware middleware.JWTConfig, supabaseDriver *infrastructure.DriverSupabase, appDriver *infrastructure.DriverAppService) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, utility.NewSuccessResponse("pong", nil))
	})

	fetchRepositorySupabase := _fetchRepositorySupabase.NewFetchRepositorySupabase(supabaseDriver)
	userRepositorySupabase := _userRepositorySupabase.NewUserRepositorySupabase(supabaseDriver)
	productRepositorySupabase := _productRepositorySupabase.NewProductRepositorySupabase(supabaseDriver)
	storeRepositorySupabase := _storeRepositorySupabase.NewStoreRepositorySupabase(supabaseDriver)
	cartRepositorySupabase := _cartRepositorySupabase.NewCartRepositorySupabase(supabaseDriver)
	invoiceRepositorySupabase := _invoiceRepositorySupabase.NewInvoiceRepositorySupabase(supabaseDriver)

	fetchUseCase := _fetchUseCase.NewFetchUseCase(fetchRepositorySupabase)
	userUseCase := _userUseCase.NewUserUseCase(userRepositorySupabase, fetchRepositorySupabase, appDriver)
	productUseCase := _productUseCase.NewProductUseCase(productRepositorySupabase, fetchRepositorySupabase)
	storeUseCase := _storeUseCase.NewStoreUseCase(storeRepositorySupabase)
	cartUseCase := _cartUseCase.NewCartUseCase(cartRepositorySupabase)
	invoiceUseCase := _invoiceUseCase.NewInvoiceUseCase(invoiceRepositorySupabase, cartRepositorySupabase, fetchRepositorySupabase)

	userGroup := e.Group("/u", middleware.JWTWithConfig(authMiddleware))

	_fetchHandler.NewFetchHandler(e, fetchUseCase)
	_userHandler.NewUserHandler(e, userGroup, userUseCase, appDriver)
	_productHandler.NewProductHandler(userGroup, productUseCase, appDriver)
	_storeHandler.NewStoreHandler(userGroup, storeUseCase, appDriver)
	_cartHandler.NewCartHandler(userGroup, cartUseCase, appDriver)
	_invoiceHandler.NewInvoiceHandler(userGroup, invoiceUseCase, appDriver)
}
