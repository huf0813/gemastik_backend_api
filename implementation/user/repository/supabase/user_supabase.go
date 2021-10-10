package supabase

import (
	"encoding/json"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"net/http"
)

type UserRepositorySupabase struct {
	DB            *infrastructure.DriverSupabase
	TableUser     string
	TableSupplier string
}

func NewUserRepositorySupabase(db *infrastructure.DriverSupabase) domain.UserRepository {
	return &UserRepositorySupabase{DB: db, TableUser: "users", TableSupplier: "suppliers"}
}

func (u *UserRepositorySupabase) SignUp(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := u.DB.RequestFormula(u.TableUser, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := u.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}

func (u *UserRepositorySupabase) ClaimSupplier(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := u.DB.RequestFormula(u.TableSupplier, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := u.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}
