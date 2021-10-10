package supabase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"io"
	"net/http"
)

type InvoiceRepositorySupabase struct {
	DB                        *infrastructure.DriverSupabase
	TableInvoice              string
	TableInvoiceProduct       string
	TableInvoiceProductReview string
}

func NewInvoiceRepositorySupabase(db *infrastructure.DriverSupabase) domain.InvoiceRepository {
	return &InvoiceRepositorySupabase{DB: db, TableInvoice: "invoices", TableInvoiceProduct: "invoice_products", TableInvoiceProductReview: "invoice_product_reviews"}
}

func (i *InvoiceRepositorySupabase) CreateInvoice(value map[string]string) (int64, string, error) {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return 0, "", err
	}

	request, err := i.DB.RequestFormula(i.TableInvoice, http.MethodPost, valueByte)
	if err != nil {
		return 0, "", err
	}

	response, err := i.DB.ExecuteRequestFormula(request)
	if err != nil {
		return 0, "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	var responseJSON []map[string]interface{}
	if err := json.Unmarshal(body, &responseJSON); err != nil {
		return 0, "", err
	}
	if len(responseJSON) <= 0 {
		return 0, "", errors.New("entry not found")
	}
	result := responseJSON[0]
	invoiceID := int64(result["id"].(float64))
	code := result["code"].(string)

	return invoiceID, code, nil
}

func (i *InvoiceRepositorySupabase) CreateInvoiceProduct(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := i.DB.RequestFormula(i.TableInvoiceProduct, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := i.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}

func (i *InvoiceRepositorySupabase) CreateInvoiceProductReview(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := i.DB.RequestFormula(i.TableInvoiceProductReview, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := i.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}

func (i *InvoiceRepositorySupabase) UpdateInvoiceStatus(value map[string]int64, userID int64, InvoiceCode string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := i.DB.RequestFormula(i.TableInvoice, http.MethodPatch, valueByte)
	if err != nil {
		return err
	}
	queries := request.URL.Query()
	queries.Add("user_id", fmt.Sprintf("eq.%d", userID))
	queries.Add("code", fmt.Sprintf("eq.%s", InvoiceCode))
	request.URL.RawQuery = queries.Encode()

	if _, err := i.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}
