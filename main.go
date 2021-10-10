package main

import (
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	_gudityMiddleware "github.com/huf0813/gemastik_api_backend_supabase/middleware"
	"github.com/huf0813/gemastik_api_backend_supabase/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	appDriver, err := infrastructure.NewDriverApp()
	if err != nil {
		panic(err)
	}

	supabaseDriver, err := infrastructure.NewDriverSupabase()
	if err != nil {
		panic(err)
	}

	gudityMiddleware, err := _gudityMiddleware.NewAuthMiddleware(appDriver.Secret)
	if err != nil {
		panic(err)
	}

	route.NewRoutes(e, gudityMiddleware, &supabaseDriver, &appDriver)

	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appDriver.Port)))
}
