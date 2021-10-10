package domain

type CreateProductRequest struct {
	Name          string `json:"name"`
	Quality       string `json:"quality"`
	Description   string `json:"description"`
	Thumbnail     string `json:"thumbnail"`
	Price         int64  `json:"price"`
	ProductTypeID int64  `json:"product_type_id"`
}

type ProductRepository interface {
	CreateProduct(value map[string]string) error
	UpdateProduct(value map[string]string, productID, supplierID int64) error
	DeleteProduct(productID, supplierID int64) error
}

type ProductUseCase interface {
	CreateProduct(create *CreateProductRequest, userID int64) error
	UpdateProduct(value *CreateProductRequest, productID, userID int64) error
	DeleteProduct(productID, userID int64) error
}
