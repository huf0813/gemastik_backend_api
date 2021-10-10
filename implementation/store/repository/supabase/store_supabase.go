package supabase

import (
	"encoding/json"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"net/http"
)

type StoreRepositorySupabase struct {
	DB         *infrastructure.DriverSupabase
	TableStore string
}

func NewStoreRepositorySupabase(db *infrastructure.DriverSupabase) domain.StoreRepository {
	return &StoreRepositorySupabase{DB: db, TableStore: "stores"}
}

func (s *StoreRepositorySupabase) ClaimStore(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := s.DB.RequestFormula(s.TableStore, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := s.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}
