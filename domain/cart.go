package domain

type AddProductToCartRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type CartRepository interface {
	AddProductToCart(value map[string]string) error
	DeleteProductFromCart(userID, productID, cartID int64) error
}

type CartUseCase interface {
	AddProductToCart(c *AddProductToCartRequest, userID int64) error
	DeleteProductFromCart(userID, productID, cartID int64) error
}
