package supabase

import (
	"encoding/json"
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"net/http"
)

type ProductRepositorySupabase struct {
	DB           *infrastructure.DriverSupabase
	TableProduct string
}

func NewProductRepositorySupabase(db *infrastructure.DriverSupabase) domain.ProductRepository {
	return &ProductRepositorySupabase{DB: db, TableProduct: "products"}
}

func (p *ProductRepositorySupabase) CreateProduct(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := p.DB.RequestFormula(p.TableProduct, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := p.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositorySupabase) UpdateProduct(value map[string]string, productID, supplierID int64) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := p.DB.RequestFormula(p.TableProduct, http.MethodPatch, valueByte)
	if err != nil {
		return err
	}
	queries := request.URL.Query()
	queries.Add("id", fmt.Sprintf("eq.%d", productID))
	queries.Add("supplier_id", fmt.Sprintf("eq.%d", supplierID))
	request.URL.RawQuery = queries.Encode()

	if _, err := p.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositorySupabase) DeleteProduct(productID, supplierID int64) error {
	request, err := p.DB.RequestFormula(p.TableProduct, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	queries := request.URL.Query()
	queries.Add("id", fmt.Sprintf("eq.%d", productID))
	queries.Add("supplier_id", fmt.Sprintf("eq.%d", supplierID))
	request.URL.RawQuery = queries.Encode()

	if _, err := p.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}
