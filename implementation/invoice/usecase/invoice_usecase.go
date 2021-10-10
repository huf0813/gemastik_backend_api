package usecase

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
)

type InvoiceUseCase struct {
	InvoiceRepositorySupabase domain.InvoiceRepository
	CartRepositorySupabase    domain.CartRepository
	FetchRepositorySupabase   domain.FetchRepository
}

func NewInvoiceUseCase(i domain.InvoiceRepository, c domain.CartRepository, f domain.FetchRepository) domain.InvoiceUseCase {
	return &InvoiceUseCase{InvoiceRepositorySupabase: i, CartRepositorySupabase: c, FetchRepositorySupabase: f}
}

func (i *InvoiceUseCase) CreateInvoice(c *domain.CreateInvoiceRequest, userID int64) (string, error) {
	createInvoice := map[string]string{
		"code":              uuid.New().String(),
		"user_id":           fmt.Sprintf("%d", userID),
		"bank_id":           fmt.Sprintf("%d", c.BankID),
		"expedition_id":     fmt.Sprintf("%d", c.ExpeditionID),
		"invoice_status_id": fmt.Sprintf("%d", 1),
	}
	invoiceID, code, err := i.InvoiceRepositorySupabase.CreateInvoice(createInvoice)
	if err != nil {
		return "", err
	}

	for _, v := range c.Carts {
		queries := map[string]string{
			"select": "id,quantity,product_id",
			"id":     fmt.Sprintf("eq.%d", v),
		}
		result, err := i.FetchRepositorySupabase.FetchValue("carts", queries)
		if err != nil {
			return "", err
		}
		if len(result) <= 0 {
			return "", errors.New("row is empty")
		}
		tempResult := result[0]
		cartID := int64(tempResult["id"].(float64))
		quantity := int64(tempResult["quantity"].(float64))
		productID := int64(tempResult["product_id"].(float64))

		createInvoiceProduct := map[string]string{
			"invoice_id": fmt.Sprintf("%d", invoiceID),
			"product_id": fmt.Sprintf("%d", productID),
			"quantity":   fmt.Sprintf("%d", quantity),
		}
		if err := i.InvoiceRepositorySupabase.CreateInvoiceProduct(createInvoiceProduct); err != nil {
			return "", err
		}

		if err := i.CartRepositorySupabase.DeleteProductFromCart(userID, productID, cartID); err != nil {
			return "", err
		}
	}

	return code, nil
}

func (i *InvoiceUseCase) UpdateInvoiceStatus(invoiceStatusID, userID int64, code string) error {
	updateInvoiceStatus := map[string]int64{
		"invoice_status_id": invoiceStatusID,
	}
	if err := i.InvoiceRepositorySupabase.UpdateInvoiceStatus(updateInvoiceStatus, userID, code); err != nil {
		return err
	}

	return nil
}

func (i *InvoiceUseCase) CreateInvoiceProductReview(c *domain.CreateInvoiceProductReviewRequest) error {
	createInvoiceProductReview := map[string]string{
		"invoice_product_id": fmt.Sprintf("%d", c.InvoiceProductID),
		"star":               fmt.Sprintf("%d", c.Star),
		"review":             c.Review,
	}
	if err := i.InvoiceRepositorySupabase.CreateInvoiceProductReview(createInvoiceProductReview); err != nil {
		return err
	}

	return nil
}
