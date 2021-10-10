package domain

type CreateInvoiceRequest struct {
	BankID       int64   `json:"bank_id"`
	ExpeditionID int64   `json:"expedition_id"`
	Carts        []int64 `json:"carts"`
}

type CreateInvoiceProductReviewRequest struct {
	InvoiceProductID int64  `json:"invoice_product_id"`
	ProductID        int64  `json:"product_id"`
	UserID           int64  `json:"user_id"`
	Star             int64  `json:"star"`
	Review           string `json:"review"`
}

type InvoiceRepository interface {
	CreateInvoice(value map[string]string) (int64, string, error)
	CreateInvoiceProduct(value map[string]string) error
	UpdateInvoiceStatus(value map[string]int64, userID int64, code string) error
	CreateInvoiceProductReview(value map[string]string) error
}

type InvoiceUseCase interface {
	CreateInvoice(c *CreateInvoiceRequest, userID int64) (string, error)
	UpdateInvoiceStatus(invoiceStatusID, userID int64, code string) error
	CreateInvoiceProductReview(c *CreateInvoiceProductReviewRequest) error
}
