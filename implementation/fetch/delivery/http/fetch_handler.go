package http

import (
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FetchHandler struct {
	FetchUseCase domain.FetchUseCase
}

func NewFetchHandler(e *echo.Echo, f domain.FetchUseCase) {
	handler := &FetchHandler{FetchUseCase: f}
	e.GET("/fetch/:table", handler.Fetch)
}

func (f *FetchHandler) Fetch(c echo.Context) error {
	tempQueries := c.QueryParams()
	queries := map[string]string{}
	for i, val := range tempQueries {
		queries[i] = val[0]
	}
	table := c.Param("table")

	result, err := f.FetchUseCase.FetchValue(table, queries)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewFailResponse(err.Error()))
	}

	msg := fmt.Sprintf("fetch data from %s successfully", table)
	return c.JSON(http.StatusOK, utility.NewSuccessResponse(msg, result))
}

