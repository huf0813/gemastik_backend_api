package supabase

import (
	"encoding/json"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"io"
	"net/http"
)

type FetchRepositorySupabase struct {
	DB *infrastructure.DriverSupabase
}

func NewFetchRepositorySupabase(db *infrastructure.DriverSupabase) domain.FetchRepository {
	return &FetchRepositorySupabase{DB: db}
}

func (f *FetchRepositorySupabase) FetchValue(table string, queryMaps map[string]string) ([]map[string]interface{}, error) {
	request, err := f.DB.RequestFormula(table, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	queries := request.URL.Query()
	for i, val := range queryMaps {
		queries.Add(i, val)
	}
	request.URL.RawQuery = queries.Encode()

	res, err := f.DB.ExecuteRequestFormula(request)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
